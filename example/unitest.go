package main

import "github.com/adesegun/presentj"

func main() {
	presentj.MavenTest("Sample", `
		package test;
		import org.junit.Test;
		import static org.assertj.core.api.Assertions.assertThat;

		public class Sample {
		    // {start OMIT
		    @Test
		    public void exampleTest() {
			assertThat("Test").hasSize(4);
		    }
		    // end} OMIT
		}
`)
}
