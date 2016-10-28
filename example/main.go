package main

import (
	"github.com/alexcrownus/presentj"
)

func main() {
	presentj.JavaExec(`
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
