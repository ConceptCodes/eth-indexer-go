// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.833
package views

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"strconv"

	"github.com/conceptcodes/eth-indexer-go/internal/models"
	"github.com/conceptcodes/eth-indexer-go/views/components"
)

func Account(data models.AccountData) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Var2 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
			templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
			templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
			if !templ_7745c5c3_IsBuffer {
				defer func() {
					templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
					if templ_7745c5c3_Err == nil {
						templ_7745c5c3_Err = templ_7745c5c3_BufErr
					}
				}()
			}
			ctx = templ.InitializeContext(ctx)
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<div class=\"bg-gray-50\"><!-- Navigation -->")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = components.Nav().Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 2, "<!-- Main Content --><main class=\"container mx-auto px-4 py-8\"><!-- Account Overview --><div class=\"bg-white rounded-lg shadow-sm border p-6 mb-6\"><div class=\"flex flex-col md:flex-row md:items-center md:justify-between gap-4\"><div><div class=\"flex items-center space-x-3 mb-2\"><h1 class=\"text-xl font-bold\">Account</h1></div><div class=\"flex items-center space-x-2\"><span class=\"font-mono text-gray-600\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var3 string
			templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(data.Address)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/account.templ`, Line: 28, Col: 39}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 3, "</span> ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templ.RenderScriptItems(ctx, templ_7745c5c3_Buffer, templ.JSFuncCall("copyToClipboard", data.Address))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 4, "<button class=\"text-gray-400 hover:text-gray-600\" onClick=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var4 templ.ComponentScript = templ.JSFuncCall("copyToClipboard", data.Address)
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var4.Call)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 5, "\"><i data-lucide=\"copy\" class=\"h-4 w-4\"></i></button></div></div><div class=\"flex flex-col items-start md:items-end\"><div class=\"text-sm text-gray-500 mb-1\">Balance</div><div class=\"text-2xl font-bold\">245.89 ETH</div><div class=\"text-gray-500\">$548,923.45 USD</div></div></div></div><!-- Stats Grid --><div class=\"grid grid-cols-1 md:grid-cols-4 gap-6 mb-6\"><div class=\"bg-white rounded-lg shadow-sm border p-4\"><div class=\"text-sm text-gray-500 mb-1\">Transactions</div><div class=\"text-2xl font-bold\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var5 string
			templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(strconv.FormatInt(data.TxCount, 10))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/account.templ`, Line: 47, Col: 84}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 6, "</div></div><div class=\"bg-white rounded-lg shadow-sm border p-4\"><div class=\"text-sm text-gray-500 mb-1\">Token Transfers</div><div class=\"text-2xl font-bold\">567</div></div><div class=\"bg-white rounded-lg shadow-sm border p-4\"><div class=\"text-sm text-gray-500 mb-1\">NFTs Owned</div><div class=\"text-2xl font-bold\">23</div></div><div class=\"bg-white rounded-lg shadow-sm border p-4\"><div class=\"text-sm text-gray-500 mb-1\">First Transaction</div><div class=\"text-2xl font-bold\">256d ago</div></div></div><!-- Tabs --><div class=\"border-b mb-6\"><nav class=\"flex space-x-6\"><button class=\"px-1 py-4 text-indigo-600 font-medium border-b-2 border-indigo-600\">Transactions</button> <button class=\"px-1 py-4 text-gray-500 font-medium hover:text-gray-700\">Tokens</button> <button class=\"px-1 py-4 text-gray-500 font-medium hover:text-gray-700\">NFTs</button> <button class=\"px-1 py-4 text-gray-500 font-medium hover:text-gray-700\">Analytics</button></nav></div><!-- Transactions Table -->")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = components.TransactionTable(data.Txs, true, true, data.PageNumber, data.TotalPages).Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 7, "<!-- Token Holdings --><div class=\"mt-6 bg-white rounded-lg shadow-sm border\"><div class=\"p-6 border-b\"><h2 class=\"text-lg font-semibold\">Token Holdings</h2></div><div class=\"overflow-x-auto\"><table class=\"w-full\"><thead><tr class=\"border-b bg-gray-50\"><th class=\"px-6 py-3 text-left text-sm font-medium text-gray-500\">Token</th><th class=\"px-6 py-3 text-left text-sm font-medium text-gray-500\">Symbol</th><th class=\"px-6 py-3 text-right text-sm font-medium text-gray-500\">Balance</th><th class=\"px-6 py-3 text-right text-sm font-medium text-gray-500\">Value</th></tr></thead> <tbody><tr class=\"border-b hover:bg-gray-50\"><td class=\"px-6 py-4\"><div class=\"flex items-center space-x-3\"><img src=\"https://via.placeholder.com/32\" class=\"w-8 h-8 rounded-full\" alt=\"Token\"> <a href=\"#\" class=\"text-indigo-600 hover:text-indigo-700\">Uniswap</a></div></td><td class=\"px-6 py-4\">UNI</td><td class=\"px-6 py-4 text-right\">1,234.56</td><td class=\"px-6 py-4 text-right\">$2,469.12</td></tr><!-- Add more token rows as needed --></tbody></table></div></div></main></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			return nil
		})
		templ_7745c5c3_Err = Page().Render(templ.WithChildren(ctx, templ_7745c5c3_Var2), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
