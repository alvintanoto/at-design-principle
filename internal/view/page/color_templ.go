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

func Colorpage(data dto.ViewBaseDTO) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div><div class=\"w-full\"><div class=\"p-4\"><div class=\"text-5xl font-semibold\">Color </div><div class=\"text-base my-2\">Our color convention pick for our site design. </div><div class=\"min-w-full min-h-[480px] rounded-md flex flex-row relative\"><div class=\"absolute w-full flex flex-col items-center m-4\"><div class=\"bg-primary w-[240px] my-1 text-white items-center text-center rounded-sm px-2 py-1\">Primary</div><div class=\"bg-success w-[240px] my-1 text-white items-center text-center rounded-sm px-2 py-1\">Success</div><div class=\"bg-warning w-[240px] my-1 text-white items-center text-center rounded-sm px-2 py-1\">Warning</div><div class=\"bg-danger w-[240px] my-1 text-white items-center text-center rounded-sm px-2 py-1\">Danger</div><div class=\"flex flex-row w-full my-1 items-center justify-around\"><div class=\"text-center text-text rounded-sm px-2 py-1\">Text </div><div class=\"text-center text-dark-text rounded-sm px-2 py-1\">Text </div></div><div class=\"flex flex-row w-full my-1 items-center justify-around\"><div class=\"text-center text-text border border-default-border rounded-sm px-4 py-1\">Border </div><div class=\"text-center text-dark-text border border-dark-default-border rounded-sm px-4 py-1\">Border </div></div></div><div class=\"bg-layout-background flex-1 rounded-l-md\"></div><div class=\"bg-dark-layout-background flex-1 rounded-r-md\"></div></div></div></div></div></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
