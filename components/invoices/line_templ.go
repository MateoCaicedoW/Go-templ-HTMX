// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.543
package invoices

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import "fmt"
import "templ/models"

func Line(id string, line models.InvoiceLine) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<tr id=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(fmt.Sprintf("line_%s", id)))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"><td class=\"px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900\"><button @click=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(fmt.Sprintf("count--; deleted=%s", id)))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-delete=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(fmt.Sprintf("/remove-line?line_%s", id)))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-target=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(fmt.Sprintf("#line_%s", id)))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-swap=\"outerHTML\" _=\"on htmx:afterRequest(e) trigger setSubtotal on #subtotal end on htmx:afterRequest(e) trigger setTotal on #total end on htmx:afterRequest(e) wait 100ms then trigger updateLines on #lines-container end\" class=\"text-red-600 hover:text-red-900 mt-2\"><svg xmlns=\"http://www.w3.org/2000/svg\" width=\"20\" height=\"20\" viewBox=\"0 0 11 13\" fill=\"none\"><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M1.73594 13C1.46094 13 1.22031 12.8917 1.01406 12.675C0.807813 12.4583 0.704687 12.2056 0.704687 11.9167V1.625H0V0.541667H3.23125V0H7.76875V0.541667H11V1.625H10.2953V11.9167C10.2953 12.2056 10.1922 12.4583 9.98594 12.675C9.77969 12.8917 9.53906 13 9.26406 13H1.73594ZM9.26406 1.625H1.73594V11.9167H9.26406V1.625ZM4.58906 10.3639H3.55781V3.15972H4.58906V10.3639ZM6.41094 10.3639H7.44219V3.15972H6.41094V10.3639Z\" fill=\"#FF3B30\"></path></svg></button></td><td class=\"px-6 py-4 whitespace-nowrap text-sm text-gray-500\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = Input(templ.Attributes{"type": "text", "class": "w-full rounded-md border p-2 mt-2 bg-slate-600 text-white border-slate-800", "value": line.Description, "name": fmt.Sprintf("InvoiceLines[%s].Description", id)}).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</td><td class=\"px-6 py-4 whitespace-nowrap text-sm text-gray-500\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = Input(templ.Attributes{
			"type":       "number",
			"class":      "w-full rounded-md border p-2 mt-2 bg-slate-600 text-white border-slate-800",
			"value":      line.UnitPriceString(),
			"name":       fmt.Sprintf("InvoiceLines[%s].UnitPrice", id),
			"hx-post":    fmt.Sprintf("/set-total-line?line=%s", id),
			"hx-target":  fmt.Sprintf("#total_line_%s", id),
			"hx-swap":    "innerHTML",
			"hx-trigger": "keyup changed delay:200ms",
			"hx-include": fmt.Sprintf("[name='InvoiceLines[%s].Quantity']", id),
			"_":          "on htmx:afterRequest(e) trigger setSubtotal on #subtotal end on htmx:afterRequest(e) trigger setTotal on #total"}).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</td><td class=\"px-6 py-4 whitespace-nowrap text-sm text-gray-500\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = Input(templ.Attributes{
			"type":       "number",
			"class":      "w-full rounded-md border p-2 mt-2 bg-slate-600 text-white border-slate-800",
			"value":      fmt.Sprintf("%d", line.Quantity),
			"name":       fmt.Sprintf("InvoiceLines[%s].Quantity", id),
			"hx-post":    fmt.Sprintf("/set-total-line?line=%s", id),
			"hx-target":  fmt.Sprintf("#total_line_%s", id),
			"hx-swap":    "innerHTML",
			"hx-trigger": "keyup changed delay:200ms",
			"hx-include": fmt.Sprintf("[name='InvoiceLines[%s].UnitPrice']", id),
			"_":          "on htmx:afterRequest(e) trigger setSubtotal on #subtotal end on htmx:afterRequest(e) trigger setTotal on #total"}).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</td><td id=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(fmt.Sprintf("total_line_%s", id)))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"px-6 py-4 whitespace-nowrap text-sm text-gray-500\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = TotalLine(line.Total()).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</td></tr>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
