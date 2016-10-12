package main

import "github.com/adesegun/presentj"

func main() {
	presentj.MavenExec("Test", `
		// {start OMIT
		package main;
		public class Test {
			public static void main(String[] args) {
				System.out.println("Hello world from Maven!");
			}
		}
		// end} OMIT
`)
}
