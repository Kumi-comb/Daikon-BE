// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "Daikon-BE": CLI Commands
//
// Command:
// $ goagen
// --design=github.com/Kumi-comb/Daikon-BE/design
// --out=$(GOPATH)\src\github.com\Kumi-comb\Daikon-BE
// --version=v1.3.1

package cli

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Kumi-comb/Daikon-BE/client"
	"github.com/goadesign/goa"
	goaclient "github.com/goadesign/goa/client"
	uuid "github.com/goadesign/goa/uuid"
	"github.com/spf13/cobra"
	"log"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

type (
	// CreateAccountAccountsCommand is the command line data structure for the createAccount action of accounts
	CreateAccountAccountsCommand struct {
		Payload     string
		ContentType string
		PrettyPrint bool
	}

	// DeviceListDevicesCommand is the command line data structure for the deviceList action of devices
	DeviceListDevicesCommand struct {
		PrettyPrint bool
	}

	// GenerateLinkCodeDevicesCommand is the command line data structure for the generateLinkCode action of devices
	GenerateLinkCodeDevicesCommand struct {
		// Device unique code
		UniqueCode  string
		PrettyPrint bool
	}

	// LinkDeviceDevicesCommand is the command line data structure for the linkDevice action of devices
	LinkDeviceDevicesCommand struct {
		// LinkCode
		LinkCode    string
		PrettyPrint bool
	}

	// SigninJWTCommand is the command line data structure for the signin action of jwt
	SigninJWTCommand struct {
		Payload     string
		ContentType string
		PrettyPrint bool
	}

	// SupportLineResourcesCommand is the command line data structure for the supportLine action of resources
	SupportLineResourcesCommand struct {
		PrettyPrint bool
	}

	// GetStatusCommand is the command line data structure for the get action of status
	GetStatusCommand struct {
		// Device unique code
		UniqueCode  string
		PrettyPrint bool
	}

	// SettingsStatusCommand is the command line data structure for the settings action of status
	SettingsStatusCommand struct {
		Payload     string
		ContentType string
		// Device unique code
		UniqueCode  string
		PrettyPrint bool
	}
)

// RegisterCommands registers the resource action CLI commands.
func RegisterCommands(app *cobra.Command, c *client.Client) {
	var command, sub *cobra.Command
	command = &cobra.Command{
		Use:   "create-account",
		Short: `アカウントの作成を行う`,
	}
	tmp1 := new(CreateAccountAccountsCommand)
	sub = &cobra.Command{
		Use:   `accounts ["/accounts"]`,
		Short: ``,
		Long: `

Payload example:

{
   "password": "Sunt voluptas.",
   "screenName": "Et a omnis laborum in molestiae."
}`,
		RunE: func(cmd *cobra.Command, args []string) error { return tmp1.Run(c, args) },
	}
	tmp1.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp1.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "device-list",
		Short: `デバイス一覧を返す`,
	}
	tmp2 := new(DeviceListDevicesCommand)
	sub = &cobra.Command{
		Use:   `devices ["/devices"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp2.Run(c, args) },
	}
	tmp2.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp2.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "generate-link-code",
		Short: `ユニークコードを使用してリンクコードを生成する`,
	}
	tmp3 := new(GenerateLinkCodeDevicesCommand)
	sub = &cobra.Command{
		Use:   `devices ["/devices/UNIQUECODE"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp3.Run(c, args) },
	}
	tmp3.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp3.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "get",
		Short: `ユニークコードを基にデバイスが本来あるべき状態のデータを取得する`,
	}
	tmp4 := new(GetStatusCommand)
	sub = &cobra.Command{
		Use:   `status ["/status/UNIQUECODE"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp4.Run(c, args) },
	}
	tmp4.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp4.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "link-device",
		Short: `リンクコードを使用してユニークコードとアカウントを鎖付けする`,
	}
	tmp5 := new(LinkDeviceDevicesCommand)
	sub = &cobra.Command{
		Use:   `devices ["/devices/LINKCODE"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp5.Run(c, args) },
	}
	tmp5.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp5.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "settings",
		Short: `デバイスの設定を定義する`,
	}
	tmp6 := new(SettingsStatusCommand)
	sub = &cobra.Command{
		Use:   `status ["/status/UNIQUECODE/settings"]`,
		Short: ``,
		Long: `

Payload example:

{
   "continueTimes": "Repudiandae sunt.",
   "lines": "Labore culpa.",
   "password": "Aperiam cum.",
   "userScreenName": "Qui voluptas unde expedita qui."
}`,
		RunE: func(cmd *cobra.Command, args []string) error { return tmp6.Run(c, args) },
	}
	tmp6.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp6.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "signin",
		Short: `ID/Passペアで認証を行う`,
	}
	tmp7 := new(SigninJWTCommand)
	sub = &cobra.Command{
		Use:   `jwt ["/jwt/signin"]`,
		Short: ``,
		Long: `

Payload example:

{
   "password": "Non ullam autem.",
   "userScreenName": "Amet veniam non."
}`,
		RunE: func(cmd *cobra.Command, args []string) error { return tmp7.Run(c, args) },
	}
	tmp7.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp7.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "support-line",
		Short: `対応している路線とその固有IDを返す`,
	}
	tmp8 := new(SupportLineResourcesCommand)
	sub = &cobra.Command{
		Use:   `resources ["/resources/supports/lines"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp8.Run(c, args) },
	}
	tmp8.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp8.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
}

func intFlagVal(name string, parsed int) *int {
	if hasFlag(name) {
		return &parsed
	}
	return nil
}

func float64FlagVal(name string, parsed float64) *float64 {
	if hasFlag(name) {
		return &parsed
	}
	return nil
}

func boolFlagVal(name string, parsed bool) *bool {
	if hasFlag(name) {
		return &parsed
	}
	return nil
}

func stringFlagVal(name string, parsed string) *string {
	if hasFlag(name) {
		return &parsed
	}
	return nil
}

func hasFlag(name string) bool {
	for _, arg := range os.Args[1:] {
		if strings.HasPrefix(arg, "--"+name) {
			return true
		}
	}
	return false
}

func jsonVal(val string) (*interface{}, error) {
	var t interface{}
	err := json.Unmarshal([]byte(val), &t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func jsonArray(ins []string) ([]interface{}, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []interface{}
	for _, id := range ins {
		val, err := jsonVal(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, val)
	}
	return vals, nil
}

func timeVal(val string) (*time.Time, error) {
	t, err := time.Parse(time.RFC3339, val)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func timeArray(ins []string) ([]time.Time, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []time.Time
	for _, id := range ins {
		val, err := timeVal(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, *val)
	}
	return vals, nil
}

func uuidVal(val string) (*uuid.UUID, error) {
	t, err := uuid.FromString(val)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func uuidArray(ins []string) ([]uuid.UUID, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []uuid.UUID
	for _, id := range ins {
		val, err := uuidVal(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, *val)
	}
	return vals, nil
}

func float64Val(val string) (*float64, error) {
	t, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func float64Array(ins []string) ([]float64, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []float64
	for _, id := range ins {
		val, err := float64Val(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, *val)
	}
	return vals, nil
}

func boolVal(val string) (*bool, error) {
	t, err := strconv.ParseBool(val)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func boolArray(ins []string) ([]bool, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []bool
	for _, id := range ins {
		val, err := boolVal(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, *val)
	}
	return vals, nil
}

// Run makes the HTTP request corresponding to the CreateAccountAccountsCommand command.
func (cmd *CreateAccountAccountsCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/accounts"
	}
	var payload client.CreateAccountAccountsPayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.CreateAccountAccounts(ctx, path, &payload, cmd.ContentType)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *CreateAccountAccountsCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request body encoded in JSON")
	cc.Flags().StringVar(&cmd.ContentType, "content", "", "Request content type override, e.g. 'application/x-www-form-urlencoded'")
}

// Run makes the HTTP request corresponding to the DeviceListDevicesCommand command.
func (cmd *DeviceListDevicesCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/devices"
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.DeviceListDevices(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *DeviceListDevicesCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
}

// Run makes the HTTP request corresponding to the GenerateLinkCodeDevicesCommand command.
func (cmd *GenerateLinkCodeDevicesCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/devices/%v", url.QueryEscape(cmd.UniqueCode))
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.GenerateLinkCodeDevices(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *GenerateLinkCodeDevicesCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var uniqueCode string
	cc.Flags().StringVar(&cmd.UniqueCode, "uniqueCode", uniqueCode, `Device unique code`)
}

// Run makes the HTTP request corresponding to the LinkDeviceDevicesCommand command.
func (cmd *LinkDeviceDevicesCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/devices/%v", url.QueryEscape(cmd.LinkCode))
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.LinkDeviceDevices(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *LinkDeviceDevicesCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var linkCode string
	cc.Flags().StringVar(&cmd.LinkCode, "linkCode", linkCode, `LinkCode`)
}

// Run makes the HTTP request corresponding to the SigninJWTCommand command.
func (cmd *SigninJWTCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/jwt/signin"
	}
	var payload client.SigninJWTPayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.SigninJWT(ctx, path, &payload, cmd.ContentType)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *SigninJWTCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request body encoded in JSON")
	cc.Flags().StringVar(&cmd.ContentType, "content", "", "Request content type override, e.g. 'application/x-www-form-urlencoded'")
}

// Run makes the HTTP request corresponding to the SupportLineResourcesCommand command.
func (cmd *SupportLineResourcesCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/resources/supports/lines"
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.SupportLineResources(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *SupportLineResourcesCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
}

// Run makes the HTTP request corresponding to the GetStatusCommand command.
func (cmd *GetStatusCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/status/%v", url.QueryEscape(cmd.UniqueCode))
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.GetStatus(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *GetStatusCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var uniqueCode string
	cc.Flags().StringVar(&cmd.UniqueCode, "uniqueCode", uniqueCode, `Device unique code`)
}

// Run makes the HTTP request corresponding to the SettingsStatusCommand command.
func (cmd *SettingsStatusCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/status/%v/settings", url.QueryEscape(cmd.UniqueCode))
	}
	var payload client.SettingsStatusPayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.SettingsStatus(ctx, path, &payload, cmd.ContentType)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *SettingsStatusCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request body encoded in JSON")
	cc.Flags().StringVar(&cmd.ContentType, "content", "", "Request content type override, e.g. 'application/x-www-form-urlencoded'")
	var uniqueCode string
	cc.Flags().StringVar(&cmd.UniqueCode, "uniqueCode", uniqueCode, `Device unique code`)
}
