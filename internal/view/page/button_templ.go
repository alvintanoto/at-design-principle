// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.543
package vpage

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import "alvintanoto.id/design-principle/internal/view/component"
import "alvintanoto.id/design-principle/internal/dto"

func selectDropdown(btnType string) templ.ComponentScript {
	return templ.ComponentScript{
		Name: `__templ_selectDropdown_764f`,
		Function: `function __templ_selectDropdown_764f(btnType){document.getElementById("btn_type_select").value = btnType
}`,
		Call:       templ.SafeScript(`__templ_selectDropdown_764f`, btnType),
		CallInline: templ.SafeScriptInline(`__templ_selectDropdown_764f`, btnType),
	}
}

func Buttonpage(data dto.ButtonViewDTO) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = vcomponent.Title(data.Name).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<body class=\"max-h-[100vh] select-none text-base overflow-hidden\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = vcomponent.HeadLayout().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"grid-cols-2 md:flex md:min-h-full\"><div class=\"hidden md:flex md:min-w-[256px]\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = vcomponent.SideNavigation(data.Name).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div><div class=\"w-full max-h-[calc(100vh-64px)] overflow-y-auto\"><div class=\"p-4\"><div class=\"text-5xl font-semibold\">Button </div><div class=\"text-base my-2\">Our buttons convention.</div><div class=\"my-1\"><form action=\"/button\" method=\"post\"><div class=\"relative w-full flex flex-col\"><div class=\"absolute w-full flex flex-row justify-end items-center h-full z-0\"><svg class=\"pointer-events-none row-start-1 col-start-1 mr-2\" height=\"12px\" width=\"12px\" viewBox=\"0 0 16 16\"><path fill-rule=\"evenodd\" d=\"M1.646 4.646a.5.5 0 0 1 .708 0L8 10.293l5.646-5.647a.5.5 0 0 1 .708.708l-6 6a.5.5 0 0 1-.708 0l-6-6a.5.5 0 0 1 0-.708\"></path></svg></div><select id=\"btn_type_select\" name=\"button_type\" onchange=\"this.form.submit()\" class=\"appearance-none w-full px-2 py-1 my-1 border border-default-border z-10 bg-transparent\"><option value=\"Primary\">Primary</option> <option value=\"Success\">Success</option> <option value=\"Warning\">Warning</option> <option value=\"Danger\">Danger</option></select></div></form></div><div class=\"min-w-full min-h-[300px] rounded-md flex flex-row relative\"><div class=\"absolute w-full flex flex-col items-center my-4\"><div class=\"flex flex-row w-full my-1 items-center justify-around\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 = []any{
			"text-center text-white rounded-md px-2 py-1 cursor-pointer shadow-sm",
			templ.KV("bg-primary hover:bg-primary/80", data.Type == "Primary"),
			templ.KV("bg-success hover:bg-success/80", data.Type == "Success"),
			templ.KV("bg-warning hover:bg-warning/80", data.Type == "Warning"),
			templ.KV("bg-danger hover:bg-danger/80", data.Type == "Danger"),
		}
		templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var2...)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ.CSSClasses(templ_7745c5c3_Var2).String()))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var3 string
		templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(data.Type)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/view/page/button.templ`, Line: 52, Col: 21}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var4 = []any{
			"text-center text-white rounded-md px-2 py-1 cursor-pointer shadow-sm",
			templ.KV("bg-primary hover:bg-primary/80", data.Type == "Primary"),
			templ.KV("bg-success hover:bg-success/80", data.Type == "Success"),
			templ.KV("bg-warning hover:bg-warning/80", data.Type == "Warning"),
			templ.KV("bg-danger hover:bg-danger/80", data.Type == "Danger"),
		}
		templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var4...)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ.CSSClasses(templ_7745c5c3_Var4).String()))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var5 string
		templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(data.Type)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/view/page/button.templ`, Line: 61, Col: 21}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div><div class=\"flex flex-row w-full my-1 items-center justify-around\"><div class=\"text-center bg-disabled text-white rounded-md px-2 py-1 cursor-default\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var6 string
		templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(data.Type)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/view/page/button.templ`, Line: 64, Col: 104}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" (Disabled) </div><div class=\"text-center bg-dark-disabled text-white rounded-md px-2 py-1 cursor-default\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var7 string
		templ_7745c5c3_Var7, templ_7745c5c3_Err = templ.JoinStringErrs(data.Type)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/view/page/button.templ`, Line: 65, Col: 109}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var7))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" (Disabled)</div></div><div class=\"flex flex-row w-full my-1 items-center justify-around\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var8 = []any{
			"text-center border border-default-border text-text rounded-md px-2 py-1 cursor-pointer shadow-sm",
			templ.KV("hover:border-primary hover:text-primary", data.Type == "Primary"),
			templ.KV("hover:border-success hover:text-success", data.Type == "Success"),
			templ.KV("hover:border-warning hover:text-warning", data.Type == "Warning"),
			templ.KV("hover:border-danger hover:text-danger", data.Type == "Danger"),
		}
		templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var8...)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ.CSSClasses(templ_7745c5c3_Var8).String()))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">Default </div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var9 = []any{
			"text-center border border-dark-default-border text-dark-text rounded-md px-2 py-1 cursor-pointer shadow-sm",
			templ.KV("hover:border-primary hover:text-primary", data.Type == "Primary"),
			templ.KV("hover:border-success hover:text-success", data.Type == "Success"),
			templ.KV("hover:border-warning hover:text-warning", data.Type == "Warning"),
			templ.KV("hover:border-danger hover:text-danger", data.Type == "Danger"),
		}
		templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var9...)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ.CSSClasses(templ_7745c5c3_Var9).String()))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">Default</div></div></div><div class=\"bg-layout-background flex-1 rounded-l-md\"></div><div class=\"bg-dark-layout-background flex-1 rounded-r-md\"></div></div></div></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = selectDropdown(data.Type).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
