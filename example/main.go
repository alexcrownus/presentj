package main

import (
	"github.com/adesegun/presentj"
)

func main() {
	presentj.JavaExec("Hello", `
// {start OMIT
package main;

public class Hello {
        public static void main(String[] args) {
                System.out.println("Hello world again!");
        }
}
// end} OMIT
`)
}
