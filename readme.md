<pre>
# To embed a manifest file as a resource, you can use the rsrc tool.
go get github.com/akavel/rsrc
rsrc -manifest test.manifest -o rsrc.syso
go build
go build -ldflags="-H windowsgui" <b># To get rid of the cmd window, instead run</b>
</pre>
 