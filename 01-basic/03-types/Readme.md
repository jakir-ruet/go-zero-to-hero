## Data Types

### 1. Basic Data Types

#### Boolean: `bool`

#### Numeric

- **Signed Integers**:

| **Type**  | **Size**    | **Range (Signed)**                                        |
|-----------|-------------|-----------------------------------------------------------|
| `int8`    | 8 bits      | -128 to 127                                               |
| `int16`   | 16 bits     | -32,768 to 32,767                                         |
| `int32`   | 32 bits     | -2,147,483,648 to 2,147,483,647                           |
| `int64`   | 64 bits     | -9,223,372,036,854,775,808 to 9,223,372,036,854,775,807   |
| `int`     | 32 or 64 bits | Depends on architecture (typically 32 or 64 bits)        |

- **Unsigned Integers**:

| **Type**  | **Size**    | **Range (Unsigned)**                                      |
|-----------|-------------|-----------------------------------------------------------|
| `uint8`   | 8 bits      | 0 to 255                                                   |
| `uint16`  | 16 bits     | 0 to 65,535                                                |
| `uint32`  | 32 bits     | 0 to 4,294,967,295                                         |
| `uint64`  | 64 bits     | 0 to 18,446,744,073,709,551,615                           |
| `uint`    | 32 or 64 bits | Depends on architecture (typically 32 or 64 bits)        |

- **Floating Point**:

| **Type**  | **Size**    | **Range (Signed)**                                        |
|-----------|-------------|-----------------------------------------------------------|
| `float32` | 32 bits     | 1.5 × 10⁻⁴⁵ to 3.4 × 10³⁸                               |
| `float64` | 64 bits     | 5.0 × 10⁻³²⁴ to 1.8 × 10³⁰⁸                             |

- **Complex Numbers**:

| **Type**     | **Size**    | **Range**                                                 |
|--------------|-------------|-----------------------------------------------------------|
| `complex64`  | 64 bits     | Real and imaginary parts are `float32`                    |
| `complex128` | 128 bits    | Real and imaginary parts are `float64`                    |

#### String: `string`

### 2. Derived Data Types

- **Arrays**: `[size]Type`
- **Slices**: `[]Type`
- **Structs**: `type StructName struct { ... }`
- **Pointers**: `*Type`

### 3. Special Data Types

- **Interface**: `type InterfaceName interface { ... }`
- **Map**: `map[KeyType]ValueType`
- **Channel**: `chan Type`
- **Error**: `type error interface { Error() string }`

### 4. Advanced Data Types

- **Function Types**: `type FunctionName func(int, int) int`
- **Empty Interface**: `interface{}`
- **Type Aliases**: `type NewType OldType`
- **Type Assertion**: `value.(type)`
- **Type Switch**: `switch v := i.(type) { ... }`

### 5. Concurrency Types

- **Goroutines**: `go functionName()`
- **Channels**: `ch := make(chan Type)`
