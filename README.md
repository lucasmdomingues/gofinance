# Installation

```go 
go get github.com/lucasmdomingues/gofinance
```

# Example

### Get only quotes for currencies and exchanges.

```go
import (
	"fmt"
	"log"
	"github.com/lucasmdomingues/gofinance"
)

func main() {
	quotations, err := gofinance.GetQuotations("KEY")
	if err != nil {
		fmt.Println(err)
		return
	}
}
```

### Obtain only CDI and SELIC rates.

```go
import (
	"fmt"
	"log"
	"github.com/lucasmdomingues/gofinance"
)

func main() {
	taxes, err := gofinance.GetTaxes("KEY")
	if err != nil {
		log.Fatal(err)
	}
}

```
### HG Finance
https://hgbrasil.com/status/finance/

### TO DO

* Obtendo dados históricos
* Personalizando a resposta
