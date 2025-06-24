// package main

// import (
// 	"fmt"
// 	"strings"
// 	"time"
// )

// // ==========================================
// // PASO 1: STRUCTS BÁSICOS
// // ==========================================

// // Libro representa un libro en la biblioteca
// type Libro struct {
// 	ID       int
// 	Titulo   string
// 	Autor    string
// 	ISBN     string
// 	Paginas  int
// 	Prestado bool
// }

// // Usuario representa a un usuario de la biblioteca
// type Usuario struct {
// 	ID       int
// 	Nombre   string
// 	Email    string
// 	Telefono string
// 	Activo   bool
// }

// // Prestamo representa el préstamo de un libro
// type Prestamo struct {
// 	ID              int
// 	LibroID         int
// 	UsuarioID       int
// 	FechaPrestamo   time.Time
// 	FechaDevolucion time.Time
// 	Devuelto        bool
// }

// // ==========================================
// // PASO 2: MÉTODOS CON RECEPTOR DE VALOR
// // (Solo para LEER información, no modifican)
// // ==========================================

// // ObtenerInfo retorna información básica del libro
// func (l Libro) ObtenerInfo() string {
// 	estado := "Disponible"
// 	if l.Prestado {
// 		estado = "Prestado"
// 	}
// 	return fmt.Sprintf("[%d] %s por %s - %s", l.ID, l.Titulo, l.Autor, estado)
// }

// // EsPrestable verifica si el libro se puede prestar
// func (l Libro) EsPrestable() bool {
// 	return !l.Prestado && l.Paginas > 0
// }

// // EsLibroGrande determina si es un libro extenso
// func (l Libro) EsLibroGrande() bool {
// 	return l.Paginas > 300
// }

// // ObtenerResumen retorna un resumen del usuario
// func (u Usuario) ObtenerResumen() string {
// 	estado := "Inactivo"
// 	if u.Activo {
// 		estado = "Activo"
// 	}
// 	return fmt.Sprintf("%s (%s) - %s", u.Nombre, u.Email, estado)
// }

// // PuedePrestar verifica si el usuario puede pedir préstamos
// func (u Usuario) PuedePrestar() bool {
// 	return u.Activo && u.Email != "" && u.Nombre != ""
// }

// // ==========================================
// // PASO 3: MÉTODOS CON RECEPTOR DE PUNTERO
// // (Para MODIFICAR el estado del struct)
// // ==========================================

// // Prestar marca el libro como prestado
// func (l *Libro) Prestar() error {
// 	if l.Prestado {
// 		return fmt.Errorf("el libro '%s' ya está prestado", l.Titulo)
// 	}
// 	if l.Paginas <= 0 {
// 		return fmt.Errorf("el libro '%s' no es válido", l.Titulo)
// 	}
// 	l.Prestado = true
// 	return nil
// }

// // Devolver marca el libro como devuelto
// func (l *Libro) Devolver() error {
// 	if !l.Prestado {
// 		return fmt.Errorf("el libro '%s' no está prestado", l.Titulo)
// 	}
// 	l.Prestado = false
// 	return nil
// }

// // ActualizarInfo permite actualizar información del libro
// func (l *Libro) ActualizarInfo(titulo, autor string, paginas int) error {
// 	if titulo == "" || autor == "" {
// 		return fmt.Errorf("título y autor no pueden estar vacíos")
// 	}
// 	if paginas <= 0 {
// 		return fmt.Errorf("número de páginas debe ser positivo")
// 	}
// 	l.Titulo = titulo
// 	l.Autor = autor
// 	l.Paginas = paginas
// 	return nil
// }

// // Activar activa la cuenta del usuario
// func (u *Usuario) Activar() {
// 	u.Activo = true
// }

// // Desactivar desactiva la cuenta del usuario
// func (u *Usuario) Desactivar() {
// 	u.Activo = false
// }

// // ActualizarContacto actualiza información de contacto
// func (u *Usuario) ActualizarContacto(email, telefono string) error {
// 	if !strings.Contains(email, "@") {
// 		return fmt.Errorf("email inválido: %s", email)
// 	}
// 	u.Email = email
// 	u.Telefono = telefono
// 	return nil
// }

// // ==========================================
// // PASO 4: STRUCT PRINCIPAL CON COMPOSICIÓN
// // ==========================================

// // Biblioteca es el struct principal que maneja todo el sistema
// type Biblioteca struct {
// 	Nombre    string
// 	Direccion string
// 	Libros    []Libro
// 	Usuarios  []Usuario
// 	Prestamos []Prestamo
// 	proximoID int
// }

// // ==========================================
// // PASO 5: MÉTODOS AVANZADOS CON LÓGICA DE NEGOCIO
// // ==========================================

// // NuevaBiblioteca es un constructor (patrón común en Go)
// func NuevaBiblioteca(nombre, direccion string) *Biblioteca {
// 	return &Biblioteca{
// 		Nombre:    nombre,
// 		Direccion: direccion,
// 		Libros:    make([]Libro, 0),
// 		Usuarios:  make([]Usuario, 0),
// 		Prestamos: make([]Prestamo, 0),
// 		proximoID: 1,
// 	}
// }

// // AgregarLibro añade un nuevo libro a la biblioteca
// func (b *Biblioteca) AgregarLibro(titulo, autor, isbn string, paginas int) (*Libro, error) {
// 	if titulo == "" || autor == "" {
// 		return nil, fmt.Errorf("título y autor son obligatorios")
// 	}

// 	// Verificar que no exista un libro con el mismo ISBN
// 	for _, libro := range b.Libros {
// 		if libro.ISBN == isbn && isbn != "" {
// 			return nil, fmt.Errorf("ya existe un libro con ISBN: %s", isbn)
// 		}
// 	}

// 	libro := Libro{
// 		ID:       b.proximoID,
// 		Titulo:   titulo,
// 		Autor:    autor,
// 		ISBN:     isbn,
// 		Paginas:  paginas,
// 		Prestado: false,
// 	}
// 	b.Libros = append(b.Libros, libro)
// 	b.proximoID++
// 	return &libro, nil
// }

// // RegistrarUsuario registra un nuevo usuario
// func (b *Biblioteca) RegistrarUsuario(nombre, email, telefono string) (*Usuario, error) {
// 	if nombre == "" || email == "" {
// 		return nil, fmt.Errorf("nombre y email son obligatorios")
// 	}
// 	if !strings.Contains(email, "@") {
// 		return nil, fmt.Errorf("email inválido: %s", email)
// 	}

// 	// Verificar que no exista un usuario con el mismo email
// 	for _, usuario := range b.Usuarios {
// 		if usuario.Email == email {
// 			return nil, fmt.Errorf("ya existe un usuario con email: %s", email)
// 		}
// 	}

// 	usuario := Usuario{
// 		ID:       b.proximoID,
// 		Nombre:   nombre,
// 		Email:    email,
// 		Telefono: telefono,
// 		Activo:   true,
// 	}
// 	b.Usuarios = append(b.Usuarios, usuario)
// 	b.proximoID++
// 	return &usuario, nil
// }

// // BuscarLibro busca un libro por ID
// func (b Biblioteca) BuscarLibro(id int) *Libro {
// 	for i, libro := range b.Libros {
// 		if libro.ID == id {
// 			return &b.Libros[i] // Retorna puntero al libro original
// 		}
// 	}
// 	return nil
// }

// // BuscarUsuario busca un usuario por ID
// func (b Biblioteca) BuscarUsuario(id int) *Usuario {
// 	for i, usuario := range b.Usuarios {
// 		if usuario.ID == id {
// 			return &b.Usuarios[i] // Retorna puntero al usuario original
// 		}
// 	}
// 	return nil
// }

// // PrestarLibro realiza el préstamo de un libro
// func (b *Biblioteca) PrestarLibro(libroID, usuarioID int) error {
// 	// Buscar libro
// 	libro := b.BuscarLibro(libroID)
// 	if libro == nil {
// 		return fmt.Errorf("libro con ID %d no encontrado", libroID)
// 	}

// 	// Buscar usuario
// 	usuario := b.BuscarUsuario(usuarioID)
// 	if usuario == nil {
// 		return fmt.Errorf("usuario con ID %d no encontrado", usuarioID)
// 	}

// 	// Validar que el usuario pueda prestar
// 	if !usuario.PuedePrestar() {
// 		return fmt.Errorf("el usuario %s no puede realizar préstamos", usuario.Nombre)
// 	}

// 	// Validar que el libro se pueda prestar
// 	if !libro.EsPrestable() {
// 		return fmt.Errorf("el libro '%s' no se puede prestar", libro.Titulo)
// 	}

// 	// Realizar el préstamo
// 	if err := libro.Prestar(); err != nil {
// 		return err
// 	}

// 	// Registrar el préstamo
// 	prestamo := Prestamo{
// 		ID:              b.proximoID,
// 		LibroID:         libroID,
// 		UsuarioID:       usuarioID,
// 		FechaPrestamo:   time.Now(),
// 		FechaDevolucion: time.Now().AddDate(0, 0, 14), // 2 semanas
// 		Devuelto:        false,
// 	}
// 	b.Prestamos = append(b.Prestamos, prestamo)
// 	b.proximoID++
// 	return nil
// }

// // DevolverLibro procesa la devolución de un libro
// func (b *Biblioteca) DevolverLibro(libroID int) error {
// 	// Buscar libro
// 	libro := b.BuscarLibro(libroID)
// 	if libro == nil {
// 		return fmt.Errorf("libro con ID %d no encontrado", libroID)
// 	}

// 	// Buscar préstamo activo
// 	var prestamoActivo *Prestamo
// 	for i := range b.Prestamos {
// 		if b.Prestamos[i].LibroID == libroID && !b.Prestamos[i].Devuelto {
// 			prestamoActivo = &b.Prestamos[i]
// 			break
// 		}
// 	}

// 	if prestamoActivo == nil {
// 		return fmt.Errorf("no se encontró préstamo activo para el libro '%s'", libro.Titulo)
// 	}

// 	// Realizar la devolución
// 	if err := libro.Devolver(); err != nil {
// 		return err
// 	}

// 	// Marcar préstamo como devuelto
// 	prestamoActivo.Devuelto = true
// 	return nil
// }

// // ObtenerEstadisticas retorna estadísticas de la biblioteca
// func (b Biblioteca) ObtenerEstadisticas() string {
// 	totalLibros := len(b.Libros)
// 	librosPrestados := 0
// 	usuariosActivos := 0
// 	prestamosActivos := 0

// 	for _, libro := range b.Libros {
// 		if libro.Prestado {
// 			librosPrestados++
// 		}
// 	}

// 	for _, usuario := range b.Usuarios {
// 		if usuario.Activo {
// 			usuariosActivos++
// 		}
// 	}

// 	for _, prestamo := range b.Prestamos {
// 		if !prestamo.Devuelto {
// 			prestamosActivos++
// 		}
// 	}

// 	return fmt.Sprintf(`📊 Estadísticas de %s:
// 📚 Total de libros: %d
// 📖 Libros prestados: %d
// 📕 Libros disponibles: %d
// 👥 Usuarios activos: %d
// 📋 Préstamos activos: %d`,
// 		b.Nombre, totalLibros, librosPrestados,
// 		totalLibros-librosPrestados, usuariosActivos, prestamosActivos)
// }

// // ListarLibrosDisponibles muestra todos los libros disponibles
// func (b Biblioteca) ListarLibrosDisponibles() {
// 	fmt.Println("\n📚 Libros Disponibles:")
// 	fmt.Println("=" + strings.Repeat("=", 50))
// 	disponibles := 0

// 	for _, libro := range b.Libros {
// 		if !libro.Prestado {
// 			fmt.Printf(" %s\n", libro.ObtenerInfo())
// 			if libro.EsLibroGrande() {
// 				fmt.Printf(" 📖 Libro extenso (%d páginas)\n", libro.Paginas)
// 			}
// 			disponibles++
// 		}
// 	}

// 	if disponibles == 0 {
// 		fmt.Println(" No hay libros disponibles")
// 	}
// }

// // ==========================================
// // FUNCIÓN PRINCIPAL DEMOSTRATIVA
// // ==========================================

// func main() {
// 	fmt.Println("🏛 SISTEMA DE BIBLIOTECA - DEMO PRÁCTICA")
// 	fmt.Println("=" + strings.Repeat("=", 50))

// 	// PASO 1: Crear biblioteca
// 	biblioteca := NuevaBiblioteca("Biblioteca Central", "Av. Principal 123")
// 	fmt.Printf("\n✅ Biblioteca creada: %s\n", biblioteca.Nombre)

// 	// PASO 2: Agregar libros
// 	fmt.Println("\n📚 Agregando libros...")
// 	libros := []struct {
// 		titulo, autor, isbn string
// 		paginas             int
// 	}{
// 		{"El Quijote", "Miguel de Cervantes", "978-84-376-0494-7", 863},
// 		{"Cien Años de Soledad", "Gabriel García Márquez", "978-84-376-0495-4", 471},
// 		{"Go Programming", "Alan Donovan", "978-0-13-419044-0", 380},
// 		{"Clean Code", "Robert Martin", "978-0-13-235088-4", 464},
// 	}

// 	for _, l := range libros {
// 		libro, err := biblioteca.AgregarLibro(l.titulo, l.autor, l.isbn, l.paginas)
// 		if err != nil {
// 			fmt.Printf("❌ Error: %v\n", err)
// 		} else {
// 			fmt.Printf(" ✅ Agregado: %s\n", libro.ObtenerInfo())
// 		}
// 	}

// 	// PASO 3: Registrar usuarios
// 	fmt.Println("\n👥 Registrando usuarios...")
// 	usuarios := []struct {
// 		nombre, email, telefono string
// 	}{
// 		{"Ana belen arca", "ana.arca.belu@email.com", "555-0101"},
// 		{"Carlos López", "carlos.lopez@email.com", "555-0102"},
// 		{"María Rodríguez", "maria.rodriguez@email.com", "555-0103"},
// 	}

// 	for _, u := range usuarios {
// 		usuario, err := biblioteca.RegistrarUsuario(u.nombre, u.email, u.telefono)
// 		if err != nil {
// 			fmt.Printf("❌ Error: %v\n", err)
// 		} else {
// 			fmt.Printf(" ✅ Registrado: %s\n", usuario.ObtenerResumen())
// 		}
// 	}

// 	// PASO 4: Realizar préstamos
// 	fmt.Println("\n📋 Realizando préstamos...")
// 	prestamos := []struct {
// 		libroID, usuarioID int
// 	}{
// 		{1, 1}, // Ana toma El Quijote
// 		{3, 2}, // Carlos toma Go Programming
// 		{2, 3}, // María toma Cien Años de Soledad
// 	}

// 	for _, p := range prestamos {
// 		err := biblioteca.PrestarLibro(p.libroID, p.usuarioID)
// 		if err != nil {
// 			fmt.Printf("❌ Error en préstamo: %v\n", err)
// 		} else {
// 			libro := biblioteca.BuscarLibro(p.libroID)
// 			usuario := biblioteca.BuscarUsuario(p.usuarioID)
// 			fmt.Printf(" ✅ %s prestó '%s'\n", usuario.Nombre, libro.Titulo)
// 		}
// 	}

// 	// PASO 5: Mostrar estado actual
// 	biblioteca.ListarLibrosDisponibles()

// 	// PASO 6: Devolver un libro
// 	fmt.Println("\n🔄 Devolviendo libro...")
// 	err := biblioteca.DevolverLibro(1) // Ana devuelve El Quijote
// 	if err != nil {
// 		fmt.Printf("❌ Error en devolución: %v\n", err)
// 	} else {
// 		fmt.Println(" ✅ El Quijote devuelto correctamente")
// 	}

// 	// PASO 7: Mostrar estadísticas finales
// 	fmt.Println("\n" + biblioteca.ObtenerEstadisticas())

// 	// PASO 8: Demostrar diferencia entre receptor de valor y puntero
// 	fmt.Println("\n🔍 DEMO: Diferencia entre receptores")
// 	fmt.Println("=" + strings.Repeat("=", 50))
// 	libro := biblioteca.BuscarLibro(4) // Clean Code
// 	fmt.Printf("Estado inicial: %s\n", libro.ObtenerInfo())

// 	// Intentar prestar (modifica el struct)
// 	err = libro.Prestar()
// 	if err != nil {
// 		fmt.Printf("❌ Error: %v\n", err)
// 	} else {
// 		fmt.Printf("Después del préstamo: %s\n", libro.ObtenerInfo())
// 	}

// 	// Verificar info (no modifica)
// 	fmt.Printf("¿Es prestable?: %v\n", libro.EsPrestable())
// 	fmt.Printf("¿Es libro grande?: %v\n", libro.EsLibroGrande())

// 	fmt.Println("\n🎯 ¡Demo completada! Los estudiantes pueden ver:")
// 	fmt.Println(" • Structs básicos y composición")
// 	fmt.Println(" • Métodos con receptor de valor (lectura)")
// 	fmt.Println(" • Métodos con receptor de puntero (modificación)")
// 	fmt.Println(" • Validaciones y manejo de errores")
// 	fmt.Println(" • Lógica de negocio completa")
// }
