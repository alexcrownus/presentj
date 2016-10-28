package main

import "github.com/alexcrownus/presentj"

func main() {
	presentj.MavenExec(`
		// {start OMIT
		package main;
		public class Test {
			public static void main(String[] args) {
				System.out.println("Hello world from Maven!");
			}
		}
		// end} OMIT
`, "example/pom.xml")
}
