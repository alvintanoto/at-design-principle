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

func Fontpage(data dto.ViewBaseDTO) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div><div class=\"w-full max-h-[calc(100vh-64px)] overflow-y-auto\"><div class=\"p-4\"><div class=\"text-5xl font-semibold my-1\">Font </div><div class=\"text-base my-2\">Our font convention pick for our site design. </div><div class=\"text-3xl font-semibold my-1 mt-4\">Font Size </div><hr class=\"text-disabled\"><div class=\"text-5xl my-2\">The quick brown fox jumps over the lazy dog</div><div class=\"text-4xl my-2\">The quick brown fox jumps over the lazy dog</div><div class=\"text-3xl my-2\">The quick brown fox jumps over the lazy dog</div><div class=\"text-2xl my-2\">The quick brown fox jumps over the lazy dog</div><div class=\"text-xl my-2\">The quick brown fox jumps over the lazy dog</div><div class=\"text-base my-2\">The quick brown fox jumps over the lazy dog</div><div class=\"text-sm my-2\">The quick brown fox jumps over the lazy dog</div><div class=\"text-3xl font-semibold my-1 mt-4\">Font Weight </div><hr class=\"text-disabled\"><div class=\"text-xl font-semibold my-2\">The quick brown fox jumps over the lazy dog</div><div class=\"text-xl font-medium my-2\">The quick brown fox jumps over the lazy dog</div><div class=\"text-xl font-normal my-2\">The quick brown fox jumps over the lazy dog</div><div class=\"text-xl font-light my-2\">The quick brown fox jumps over the lazy dog</div></div></div></div></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
