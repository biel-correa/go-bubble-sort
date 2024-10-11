# Entendendo o Algoritmo

Este código em Go (Golang) implementa o algoritmo de ordenação *Bubble Sort* para ordenar um array de inteiros e utiliza a biblioteca `math/rand` para gerar números aleatórios dentro do array. Vou detalhar o funcionamento por partes:

### 1. **Pacotes Importados**
```go
import (
    "fmt"
    "math/rand/v2"
)
```
- `fmt`: Usado para exibir saídas no console.
- `math/rand/v2`: Usado para gerar números inteiros aleatórios.

### 2. **Função `BubbleSort`**
```go
func BubbleSort(array []int) []int {
    for i := 0; i < len(array)-1; i++ {
        for j := 0; j < len(array)-i-1; j++ {
            if array[j] > array[j+1] {
                array[j], array[j+1] = array[j+1], array[j]
            }
        }
    }
    return array
}
```
- **Entrada**: Um array de inteiros (`[]int`).
- **Lógica**: O algoritmo *Bubble Sort* percorre o array várias vezes, comparando pares de elementos adjacentes e trocando-os se estiverem na ordem errada.
  - A cada passada pelo array (controlada pela variável `i`), o maior valor "flutua" para o final da lista.
  - O número de iterações internas (`j`) diminui progressivamente, pois os maiores elementos já estão ordenados no final.
  - A condição `if array[j] > array[j+1]` verifica se o valor na posição `j` é maior que o na posição `j+1`. Se for, eles trocam de posição.
- **Saída**: Retorna o array ordenado.

### 3. **Função `main`**
```go
func main() {
    array := make([]int, 10)
    for i := range array {
        array[i] = rand.IntN(100)
    }
    fmt.Println("Unsorted array:", array)
    fmt.Println("Sorted array:", BubbleSort(array))
}
```
- **Criação do array**: 
  - `array := make([]int, 10)` cria um array de inteiros com 10 posições.
  - O `for` preenche o array com números aleatórios de 0 a 99 gerados pela função `rand.IntN(100)`.
- **Impressão do array original**: 
  - `fmt.Println("Unsorted array:", array)` exibe o array não ordenado no console.
- **Ordenação e impressão**:
  - A função `BubbleSort(array)` é chamada para ordenar o array.
  - O array ordenado é impresso com `fmt.Println("Sorted array:", BubbleSort(array))`.

### Exemplo de saída:
- **Array antes de ser ordenado**:
  ```
  Unsorted array: [42 15 77 31 89 22 3 58 6 90]
  ```
- **Array depois de ser ordenado**:
  ```
  Sorted array: [3 6 15 22 31 42 58 77 89 90]
  ```

### Resumo:
- O código cria um array de 10 números aleatórios entre 0 e 99.
- Ele usa o algoritmo de ordenação *Bubble Sort* para ordenar o array.
- O array original e o array ordenado são impressos no console.