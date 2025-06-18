Quando Usar Cada Tipo de Return
Return Simples: Quando a função não pode falhar e não precisa retornar um erro.
Return com Erro: Quando a função pode falhar e você precisa retornar um erro.
Return nil, err: Quando a função retorna um ponteiro ou interface e pode falhar.
Return nil, nil: Quando a função retorna um ponteiro ou interface e pode não encontrar um valor válido, mas isso não é um erro.
Return com Múltiplos Valores: Quando a função retorna múltiplos valores, incluindo erros.
Return com Erro Personalizado: Quando você quer retornar um erro mais informativo ou personalizado.
Return com Erro de Contexto: Quando você quer incluir contexto adicional no erro.
Return com Erro de Wrapping: Quando você quer "embrulhar" um erro para incluir mais contexto.
Return com Erro de Log: Quando você quer registrar um erro antes de retorná-lo.


1. Return Simples
Usado quando você precisa retornar um valor de uma função sem erros.
Exemplo:

func add(a, b int) int {
    return a + b
}

2. Return com Erro
Usado quando uma função pode falhar e você precisa retornar um erro junto com um valor.
Exemplo:

func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("division by zero")
    }
    return a / b, nil
}

3. Return nil, err
Usado quando a função retorna um ponteiro ou interface e um erro. Se ocorrer um erro, você retorna nil e o erro.
Exemplo:

func getUser(id int) (*User, error) {
    // Simulando uma busca no banco de dados
    if id <= 0 {
        return nil, fmt.Errorf("invalid user ID")
    }
    user := &User{Name: "John Doe", Email: "john@example.com"}
    return user, nil
}

4. Return nil, nil
Usado quando a função retorna um ponteiro ou interface e um erro, mas não há erro e o valor retornado é nil.
Exemplo:

func findUser(id int) (*User, error) {
    // Simulando uma busca no banco de dados
    if id == 1 {
        user := &User{Name: "John Doe", Email: "john@example.com"}
        return user, nil
    }
    return nil, nil // Não encontrou o usuário, mas não é um erro
}

5. Return com Múltiplos Valores
Usado quando a função retorna múltiplos valores, incluindo erros.
Exemplo:

func getUserDetails(id int) (string, string, error) {
    if id <= 0 {
        return "", "", fmt.Errorf("invalid user ID")
    }
    name := "John Doe"
    email := "john@example.com"
    return name, email, nil
}

6. Return com Erro Personalizado
Usado quando você quer retornar um erro mais informativo ou personalizado.
Exemplo:

type UserNotFoundError struct {
    ID int
}

func (e *UserNotFoundError) Error() string {
    return fmt.Sprintf("user with ID %d not found", e.ID)
}

func getUser(id int) (*User, error) {
    if id <= 0 {
        return nil, fmt.Errorf("invalid user ID")
    }
    // Simulando uma busca no banco de dados
    if id == 1 {
        user := &User{Name: "John Doe", Email: "john@example.com"}
        return user, nil
    }
    return nil, &UserNotFoundError{ID: id}
}

7. Return com Erro de Contexto
Usado quando você quer incluir contexto adicional no erro.
Exemplo:

func getUser(id int) (*User, error) {
    if id <= 0 {
        return nil, fmt.Errorf("invalid user ID: %d", id)
    }
    // Simulando uma busca no banco de dados
    if id == 1 {
        user := &User{Name: "John Doe", Email: "john@example.com"}
        return user, nil
    }
    return nil, fmt.Errorf("user with ID %d not found", id)
}

8. Return com Erro de Wrapping
Usado quando você quer "embrulhar" um erro para incluir mais contexto.
Exemplo:

func getUser(id int) (*User, error) {
    if id <= 0 {
        return nil, fmt.Errorf("invalid user ID: %d", id)
    }
    // Simulando uma busca no banco de dados
    if id == 1 {
        user := &User{Name: "John Doe", Email: "john@example.com"}
        return user, nil
    }
    return nil, fmt.Errorf("failed to get user with ID %d: %w", id, errors.New("user not found"))
}

9. Return com Erro de Log
Usado quando você quer registrar um erro antes de retorná-lo.
Exemplo:

func getUser(id int) (*User, error) {
    if id <= 0 {
        log.Printf("Invalid user ID: %d", id)
        return nil, fmt.Errorf("invalid user ID: %d", id)
    }
    // Simulando uma busca no banco de dados
    if id == 1 {
        user := &User{Name: "John Doe", Email: "john@example.com"}
        return user, nil
    }
    log.Printf("User with ID %d not found", id)
    return nil, fmt.Errorf("user with ID %d not found", id)
}