package main

import "github.com/alexcrownus/presentj"

func main() {
	presentj.MavenTest(`
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
`, "example/pom.xml")
}
