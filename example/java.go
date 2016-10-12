package main

import "github.com/adesegun/presentj"

func main() {
	presentj.JavaExec("Test", `
		// {start OMIT
		package main;
		public class Test {
			public static void main(String[] args) {
				System.out.println("Hello world from Java!");
			}
		}
		// end} OMIT
`)
}
