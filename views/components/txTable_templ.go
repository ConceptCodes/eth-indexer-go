// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.833
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"github.com/conceptcodes/eth-indexer-go/internal/helpers"
	"github.com/conceptcodes/eth-indexer-go/internal/models"
	"strconv"
)

func TransactionTable(txs []models.SimpleTransaction, showBlockNumber, pagination bool) templ.Component {
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
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<div class=\"overflow-x-auto\"><table class=\"w-full\"><thead><tr class=\"border-b\"><th class=\"px-6 py-3 text-left text-sm font-medium text-gray-500\">Transaction Hash</th>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if showBlockNumber {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 2, "<th class=\"px-6 py-3 text-left text-sm font-medium text-gray-500\">Block</th>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 3, "<th class=\"px-6 py-3 text-left text-sm font-medium text-gray-500\">From</th><th class=\"px-6 py-3 text-left text-sm font-medium text-gray-500\">To</th><th class=\"px-6 py-3 text-right text-sm font-medium text-gray-500\">Value</th></tr></thead> <tbody>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for _, tx := range txs {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 4, "<tr class=\"border-b hover:bg-gray-50\"><td class=\"px-6 py-4\"><a href=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var2 templ.SafeURL = templ.SafeURL("/tx/" + tx.Hash)
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(string(templ_7745c5c3_Var2)))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 5, "\" class=\"text-indigo-600 hover:text-indigo-700 font-mono\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var3 string
			templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(helpers.TruncateHash(tx.Hash))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/components/txTable.templ`, Line: 28, Col: 50}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 6, "</a></td>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if showBlockNumber {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 7, "<td class=\"px-6 py-4\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var4 string
				templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(strconv.FormatUint(tx.BlockNumber, 10))
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/components/txTable.templ`, Line: 32, Col: 79}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 8, "</td>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 9, "<td class=\"px-6 py-4\"><a href=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var5 templ.SafeURL = templ.SafeURL("/address/" + tx.From)
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(string(templ_7745c5c3_Var5)))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 10, "\" class=\"text-indigo-600 hover:text-indigo-700 font-mono\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var6 string
			templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(helpers.TruncateHash(tx.From))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/components/txTable.templ`, Line: 36, Col: 50}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 11, "</a></td><td class=\"px-6 py-4\"><a href=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var7 templ.SafeURL = templ.SafeURL("/address/" + tx.To)
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(string(templ_7745c5c3_Var7)))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 12, "\" class=\"text-indigo-600 hover:text-indigo-700 font-mono\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var8 string
			templ_7745c5c3_Var8, templ_7745c5c3_Err = templ.JoinStringErrs(helpers.TruncateHash(tx.To))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/components/txTable.templ`, Line: 41, Col: 48}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var8))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 13, "</a></td><td class=\"px-6 py-4 text-right\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var9 string
			templ_7745c5c3_Var9, templ_7745c5c3_Err = templ.JoinStringErrs(helpers.FormatEthValue(tx.Value))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/components/txTable.templ`, Line: 45, Col: 50}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var9))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 14, " ETH</td></tr>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 15, "</tbody></table></div><!-- Pagination -->")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if pagination {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 16, "<div class=\"mt-6 flex flex-col sm:flex-row items-center justify-between gap-4 px-6 pb-6\"><!-- Items per page --><div class=\"flex items-center space-x-2\"><span class=\"text-sm text-gray-500\">Show</span> <select class=\"border rounded px-2 py-1 text-sm\"><option>25</option> <option>50</option> <option>100</option></select> <span class=\"text-sm text-gray-500\">entries</span></div><!-- Page info --><div class=\"text-sm text-gray-500\">Showing 1 to 25 of 150 transactions</div><!-- Pagination controls --><div class=\"flex items-center space-x-2\"><button class=\"px-3 py-1 border rounded hover:bg-gray-50 text-gray-500 flex items-center space-x-1 disabled:opacity-50 disabled:cursor-not-allowed\" disabled><i data-lucide=\"chevrons-left\" class=\"h-4 w-4\"></i> <span class=\"hidden sm:inline\">First</span></button> <button class=\"px-3 py-1 border rounded hover:bg-gray-50 text-gray-500 flex items-center space-x-1 disabled:opacity-50 disabled:cursor-not-allowed\" disabled><i data-lucide=\"chevron-left\" class=\"h-4 w-4\"></i> <span class=\"hidden sm:inline\">Previous</span></button><div class=\"flex items-center space-x-1\"><button class=\"px-3 py-1 border rounded bg-indigo-50 text-indigo-600 font-medium\">1</button> <button class=\"px-3 py-1 border rounded hover:bg-gray-50\">2</button> <button class=\"px-3 py-1 border rounded hover:bg-gray-50\">3</button> <span class=\"px-2\">...</span> <button class=\"px-3 py-1 border rounded hover:bg-gray-50\">6</button></div><button class=\"px-3 py-1 border rounded hover:bg-gray-50 text-gray-500 flex items-center space-x-1\"><span class=\"hidden sm:inline\">Next</span> <i data-lucide=\"chevron-right\" class=\"h-4 w-4\"></i></button> <button class=\"px-3 py-1 border rounded hover:bg-gray-50 text-gray-500 flex items-center space-x-1\"><span class=\"hidden sm:inline\">Last</span> <i data-lucide=\"chevrons-right\" class=\"h-4 w-4\"></i></button></div></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
