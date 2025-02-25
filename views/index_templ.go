// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.833
package views

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func Index() templ.Component {
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
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<!-- Navigation --> <nav class=\"fixed w-full bg-white/80 backdrop-blur-sm z-50 border-b\"><div class=\"container mx-auto px-4\"><div class=\"flex items-center justify-between h-16\"><a href=\"#\" class=\"text-xl font-bold\">Logo</a><div class=\"hidden md:flex space-x-8\"><a href=\"#features\" class=\"text-gray-600 hover:text-gray-900\">Features</a> <a href=\"#how-it-works\" class=\"text-gray-600 hover:text-gray-900\">How it Works</a> <a href=\"#testimonials\" class=\"text-gray-600 hover:text-gray-900\">Testimonials</a></div><button class=\"bg-indigo-600 text-white px-6 py-2 rounded-full hover:bg-indigo-700 transition-colors\">Get Started</button></div></div></nav><!-- Hero Section --> <section class=\"pt-32 pb-20 gradient-bg\"><div class=\"container mx-auto px-4\"><div class=\"flex flex-col md:flex-row items-center\"><div class=\"md:w-1/2 text-center md:text-left\"><h1 class=\"text-4xl md:text-6xl font-bold text-white mb-6\">Build Something Amazing Today</h1><p class=\"text-lg text-indigo-100 mb-8\">Launch your project with our powerful platform. Fast, secure, and scalable solutions for modern web applications.</p><div class=\"flex flex-col sm:flex-row gap-4 justify-center md:justify-start\"><a href=\"/login\" class=\"bg-white text-indigo-600 px-8 py-3 rounded-full font-medium hover:bg-gray-100 transition-colors\">Get Started</a> <button class=\"border border-white text-white px-8 py-3 rounded-full font-medium hover:bg-white/10 transition-colors\">Learn More</button></div></div><div class=\"md:w-1/2 mt-12 md:mt-0\"><img src=\"https://via.placeholder.com/600x400\" alt=\"Hero Image\" class=\"rounded-lg shadow-2xl\"></div></div></div></section><!-- Features Section --> <section id=\"features\" class=\"py-20\"><div class=\"container mx-auto px-4\"><div class=\"text-center mb-16\"><h2 class=\"text-3xl font-bold mb-4\">Why Choose Us</h2><p class=\"text-gray-600 max-w-2xl mx-auto\">We provide the tools and features you need to build exceptional products.</p></div><div class=\"grid md:grid-cols-3 gap-8\"><div class=\"p-6 border rounded-xl hover:shadow-lg transition-shadow\"><div class=\"w-12 h-12 bg-indigo-100 rounded-lg flex items-center justify-center mb-4\"><i data-lucide=\"zap\" class=\"w-6 h-6 text-indigo-600\"></i></div><h3 class=\"text-xl font-semibold mb-2\">Lightning Fast</h3><p class=\"text-gray-600\">Optimized for speed and performance, ensuring your applications run smoothly.</p></div><div class=\"p-6 border rounded-xl hover:shadow-lg transition-shadow\"><div class=\"w-12 h-12 bg-indigo-100 rounded-lg flex items-center justify-center mb-4\"><i data-lucide=\"shield\" class=\"w-6 h-6 text-indigo-600\"></i></div><h3 class=\"text-xl font-semibold mb-2\">Secure by Default</h3><p class=\"text-gray-600\">Built-in security features to protect your data and users.</p></div><div class=\"p-6 border rounded-xl hover:shadow-lg transition-shadow\"><div class=\"w-12 h-12 bg-indigo-100 rounded-lg flex items-center justify-center mb-4\"><i data-lucide=\"settings\" class=\"w-6 h-6 text-indigo-600\"></i></div><h3 class=\"text-xl font-semibold mb-2\">Easy to Use</h3><p class=\"text-gray-600\">Intuitive interface and comprehensive documentation for quick setup.</p></div></div></div></section><!-- How it Works --> <section id=\"how-it-works\" class=\"py-20 bg-gray-50\"><div class=\"container mx-auto px-4\"><div class=\"text-center mb-16\"><h2 class=\"text-3xl font-bold mb-4\">How It Works</h2><p class=\"text-gray-600 max-w-2xl mx-auto\">Get started in just a few simple steps</p></div><div class=\"grid md:grid-cols-4 gap-8\"><div class=\"text-center\"><div class=\"w-16 h-16 bg-indigo-600 rounded-full flex items-center justify-center text-white text-xl font-bold mx-auto mb-4\">1</div><h3 class=\"font-semibold mb-2\">Sign Up</h3><p class=\"text-gray-600\">Create your account in seconds</p></div><div class=\"text-center\"><div class=\"w-16 h-16 bg-indigo-600 rounded-full flex items-center justify-center text-white text-xl font-bold mx-auto mb-4\">2</div><h3 class=\"font-semibold mb-2\">Configure</h3><p class=\"text-gray-600\">Set up your preferences</p></div><div class=\"text-center\"><div class=\"w-16 h-16 bg-indigo-600 rounded-full flex items-center justify-center text-white text-xl font-bold mx-auto mb-4\">3</div><h3 class=\"font-semibold mb-2\">Deploy</h3><p class=\"text-gray-600\">Launch your application</p></div><div class=\"text-center\"><div class=\"w-16 h-16 bg-indigo-600 rounded-full flex items-center justify-center text-white text-xl font-bold mx-auto mb-4\">4</div><h3 class=\"font-semibold mb-2\">Scale</h3><p class=\"text-gray-600\">Grow with your needs</p></div></div></div></section><!-- Testimonials --> <section id=\"testimonials\" class=\"py-20\"><div class=\"container mx-auto px-4\"><div class=\"text-center mb-16\"><h2 class=\"text-3xl font-bold mb-4\">What Our Customers Say</h2><p class=\"text-gray-600 max-w-2xl mx-auto\">Don't just take our word for it</p></div><div class=\"grid md:grid-cols-3 gap-8\"><div class=\"p-6 bg-white rounded-xl shadow-lg\"><div class=\"flex items-center mb-4\"><img src=\"https://via.placeholder.com/60\" alt=\"User\" class=\"w-12 h-12 rounded-full\"><div class=\"ml-4\"><h4 class=\"font-semibold\">Sarah Johnson</h4><p class=\"text-gray-600\">CEO, TechCorp</p></div></div><p class=\"text-gray-600\">\"This platform has transformed how we build and deploy applications. Couldn't be happier with the results.\"</p></div><div class=\"p-6 bg-white rounded-xl shadow-lg\"><div class=\"flex items-center mb-4\"><img src=\"https://via.placeholder.com/60\" alt=\"User\" class=\"w-12 h-12 rounded-full\"><div class=\"ml-4\"><h4 class=\"font-semibold\">Mark Wilson</h4><p class=\"text-gray-600\">Lead Developer</p></div></div><p class=\"text-gray-600\">\"The ease of use and powerful features make this the perfect solution for our team's needs.\"</p></div><div class=\"p-6 bg-white rounded-xl shadow-lg\"><div class=\"flex items-center mb-4\"><img src=\"https://via.placeholder.com/60\" alt=\"User\" class=\"w-12 h-12 rounded-full\"><div class=\"ml-4\"><h4 class=\"font-semibold\">Emily Chen</h4><p class=\"text-gray-600\">Startup Founder</p></div></div><p class=\"text-gray-600\">\"Outstanding support and incredible performance. This platform exceeds all expectations.\"</p></div></div></div></section><!-- Newsletter --> <section class=\"py-20 gradient-bg\"><div class=\"container mx-auto px-4\"><div class=\"max-w-2xl mx-auto text-center\"><h2 class=\"text-3xl font-bold text-white mb-4\">Stay Updated</h2><p class=\"text-indigo-100 mb-8\">Subscribe to our newsletter for the latest updates and features.</p><form class=\"flex flex-col sm:flex-row gap-4 justify-center\"><input type=\"email\" placeholder=\"Enter your email\" class=\"px-6 py-3 rounded-full focus:outline-none focus:ring-2 focus:ring-white/50 flex-1 max-w-md\"> <button class=\"bg-white text-indigo-600 px-8 py-3 rounded-full font-medium hover:bg-gray-100 transition-colors\">Subscribe</button></form></div></div></section><!-- Footer --> <footer class=\"bg-gray-900 text-white py-12\"><div class=\"container mx-auto px-4\"><div class=\"grid md:grid-cols-4 gap-8\"><div><h3 class=\"text-lg font-semibold mb-4\">About</h3><p class=\"text-gray-400\">Building the future of web development with powerful, easy-to-use tools.</p></div><div><h3 class=\"text-lg font-semibold mb-4\">Quick Links</h3><ul class=\"space-y-2\"><li><a href=\"#\" class=\"text-gray-400 hover:text-white\">Home</a></li><li><a href=\"#\" class=\"text-gray-400 hover:text-white\">Features</a></li><li><a href=\"#\" class=\"text-gray-400 hover:text-white\">Pricing</a></li><li><a href=\"#\" class=\"text-gray-400 hover:text-white\">Contact</a></li></ul></div><div><h3 class=\"text-lg font-semibold mb-4\">Resources</h3><ul class=\"space-y-2\"><li><a href=\"#\" class=\"text-gray-400 hover:text-white\">Documentation</a></li><li><a href=\"#\" class=\"text-gray-400 hover:text-white\">Blog</a></li><li><a href=\"#\" class=\"text-gray-400 hover:text-white\">Support</a></li></ul></div><div><h3 class=\"text-lg font-semibold mb-4\">Connect</h3><div class=\"flex space-x-4\"><a href=\"#\" class=\"text-gray-400 hover:text-white\"><i data-lucide=\"twitter\" class=\"w-6 h-6\"></i></a> <a href=\"#\" class=\"text-gray-400 hover:text-white\"><i data-lucide=\"github\" class=\"w-6 h-6\"></i></a> <a href=\"#\" class=\"text-gray-400 hover:text-white\"><i data-lucide=\"linkedin\" class=\"w-6 h-6\"></i></a></div></div></div><div class=\"border-t border-gray-800 mt-12 pt-8 text-center text-gray-400\"><p>© <script>document.write(new Date().getFullYear())</script>Your Company. All rights reserved.</p></div></div></footer>")
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
