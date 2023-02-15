package codes

import (
	"context"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/s1ovac/gdoc/internal/config"
	"github.com/s1ovac/gdoc/internal/handlers"
	"github.com/s1ovac/gdoc/internal/middleware"
	"github.com/s1ovac/gdoc/internal/parser"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
	"io/ioutil"
	"net/http"
)

type handler struct {
	ctx context.Context
	cfg config.SheetConfig
}

func NewHandler(ctx context.Context, cfg config.SheetConfig) handlers.Handler {
	return &handler{
		ctx: ctx,
		cfg: cfg,
	}
}

func (h *handler) Register(router *httprouter.Router) {
	router.HandlerFunc(http.MethodPost, "/api_update", middleware.Middleware(h.UpdateCodes))
}

func (h *handler) UpdateCodes(w http.ResponseWriter, r *http.Request) error {
	prs := parser.NewParseHTML(h.cfg)
	data, err := prs.ParseTable(h.ctx)

	if err != nil {
		return err
	}
	credBytes, err := ioutil.ReadFile(h.cfg.CredPath())
	if err != nil {
		return err
	}
	config, err := google.JWTConfigFromJSON(credBytes, "https://www.googleapis.com/auth/spreadsheets")
	if err != nil {
		return err
	}
	client := config.Client(h.ctx)

	srv, err := sheets.NewService(h.ctx, option.WithHTTPClient(client))
	if err != nil {
		return err
	}
	// Convert sheet ID to sheet name.
	response1, err := srv.Spreadsheets.Get(h.cfg.SpreadSheetID()).Fields("sheets(properties(sheetId,title))").Do()
	if err != nil || response1.HTTPStatusCode != 200 {
		return err
	}
	sheetName := ""
	for _, v := range response1.Sheets {
		prop := v.Properties
		if prop.SheetId == int64(h.cfg.SheetID()) {
			sheetName = prop.Title
			break
		}
	}
	// Build the value range to write to the sheet
	var vr sheets.ValueRange
	values := make([][]interface{}, len(data))
	for i, d := range data {
		code := d.Code()
		description := d.Description()
		values[i] = []interface{}{code, description}
	}
	vr.Values = values

	writeRange := fmt.Sprintf("%s!A1:B%d", sheetName, len(data))
	_, err = srv.Spreadsheets.Values.Update(h.cfg.SpreadSheetID(), writeRange, &vr).ValueInputOption("USER_ENTERED").Do()
	if err != nil {
		return fmt.Errorf("failed to write to sheet %s in spreadsheet %s: %v", sheetName, h.cfg.SpreadSheetID(), err)
	}
	return nil
}
