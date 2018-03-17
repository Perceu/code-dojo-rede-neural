package main

import "testing"

func TestCaixaEletronico(t *testing.T) {
    // <setup code>
    t.Run("result10", func(t *testing.T) {
	    notas := Saque(10)
	    if !(notas[0] == 1 && notas[1] == 0 && notas[2] == 0 && notas[3] == 0){
	       t.Errorf("Incorrect value")
	    }        
    })

    t.Run("result20", func(t *testing.T) {
	    notas := Saque(20)
	    if !(notas[0] == 0 && notas[1] == 1 && notas[2] == 0 && notas[3] == 0){
	       t.Errorf("incorrect, got: %d, want: %d.", notas, 20)
	    }        
    })

    t.Run("result50", func(t *testing.T) {
	    notas := Saque(50)
	    if !(notas[0] == 0 && notas[1] == 0 && notas[2] == 1 && notas[3] == 0) {
	       t.Errorf("incorrect, got: %d, want: %d.", notas, 50)
	    }        
    })

    t.Run("result100", func(t *testing.T) {
	    notas := Saque(100)
	    if !(notas[0] == 0 && notas[1] == 0 && notas[2] == 0 && notas[3] == 1){
	       t.Errorf("incorrect, got: %d, want: %d.", notas, 100)
	    }        
    })


    t.Run("result30", func(t *testing.T) {
	    notas := Saque(30)
	    if !(notas[0] == 1 && notas[1] == 1 && notas[2] == 0 && notas[3] == 0){
	       t.Errorf("incorrect, got: %d, want: %d.", notas, 30)
	    }        
    })

    t.Run("result80", func(t *testing.T) {
	    notas := Saque(80)
	    if !(notas[0] == 1 && notas[1] == 1 && notas[2] == 1 && notas[3] == 0){
	       t.Errorf("incorrect, got: %d, want: %d.", notas, 80)
	    }        
    })

    t.Run("result0", func(t *testing.T) {
	    notas := Saque(0)
	    if !(notas[0] == 0 && notas[1] == 0 && notas[2] == 0 && notas[3] == 0){
	       t.Errorf("incorrect, got: %d, want: %d.", notas, 0)
	    }        
    })

    t.Run("result480", func(t *testing.T) {
	    notas := Saque(480)
	    if !(notas[0] == 1 && notas[1] == 1 && notas[2] == 1 && notas[3] == 4){
	       t.Errorf("incorrect, got: %d, want: %d.", notas, 480)
	    }        
    })

    t.Run("result485", func(t *testing.T) {
	    notas := Saque(485)
	    if !(notas[0] == -1 && notas[1] == -1 && notas[2] == -1 && notas[3] == -1){
	       t.Errorf("incorrect, got: %d, want: %d.", notas, 485)
	    }        
    })

}
