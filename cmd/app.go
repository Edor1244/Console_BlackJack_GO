package cmd

import (
	"Pruebas_D_GO/src/game/controllers"
	"fmt"
)

type app struct {
    deckCtrl     *controllers.deckController // Usar punteros a los controladores
    playerCtrl   *controllers.playerController
    crupierCtrl  *controllers.crupierController
    handCtrl     *controllers.handController // Si decides usar un handController
}

func (a *app) Run() {
    // ... (Lógica del juego aquí)

}

type App interface {
	Run()
}

func NewApp() App{
	a := &app{}
    a.deckCtrl = controllers.NewDeckController()
    a.handCtrl = controllers.NewHandController()
    a.playerCtrl = controllers.NewPlayerController(a.deckCtrl, a.handCtrl)
    a.crupierCtrl = controllers.NewCrupierController(a.deckCtrl, a.handCtrl)
    return a
}

func (a *app) iniciarJuego() {
    a.deckCtrl.Shuffle() // Barajar el mazo

    // Repartir cartas iniciales
    for i := 0; i < 2; i++ {
        a.playerCtrl.addCard(a.deckCtrl.Deal())
        a.crupierCtrl.addCard(a.deckCtrl.Deal())
    }

    // Calcular puntuaciones iniciales
    playerScore := a.playerCtrl.handValue()

    fmt.Println("Cartas del jugador:")
    fmt.Println(a.playerCtrl.showHand())
    fmt.Println("Score del jugador:", playerScore)

    fmt.Println("Cartas del crupier:")
    fmt.Println("Carta oculta")
    fmt.Println(a.crupierCtrl.showHand()[1]) // Mostrar solo la segunda carta

    // Verificar si hay Blackjack (opcional, si tienes el método isBlackjack)
    if a.playerCtrl.isBlackjack() {
        fmt.Println("¡Blackjack!")
		a.finalizarJuegoBJP()
    } else{
		
	}
}

func (a *app) finalizarJuegoBJP() {
    fmt.Println("\n--- Resultados Finales ---")
    fmt.Println("Mano del jugador:", a.playerCtrl.showHand(), "- Puntuación:", a.playerCtrl.handValue())
    fmt.Println("\n¡El jugador gana con Blackjack!") // Mensaje específico para este caso
}



