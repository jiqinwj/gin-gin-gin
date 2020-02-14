import (
	"github.com/gin-gonic/gin"
	 "io"
      "log"
      "os"
	{{range . }}
	 "{{.}}"
	{{ end }}

)