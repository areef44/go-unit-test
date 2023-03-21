package helper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"runtime"
	"testing"
)

func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "Muhammad",
			request: "Muhammad",
		},
		{
			name:    "Arif",
			request: "Arif",
		},
		{
			name:    "MuhammadArifBukittinggi",
			request: "MuhammadArifBukittinggi",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}
}

func BenchmarkSub(b *testing.B) {
	b.Run("Arif", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Arif")
		}
	})
	b.Run("Muhammad", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Muhammad")
		}
	})
}

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Arif")
	}
}

func BenchmarkHelloWorldMuhammad(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Muhammad")
	}
}

func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Arif",
			request:  "Arif",
			expected: "Hello Arif",
		},
		{
			name:     "Muhammad",
			request:  "Muhammad",
			expected: "Hello Muhammad",
		},
		{
			name:     "Bukittinggi",
			request:  "Bukittinggi",
			expected: "Hello Bukittinggi",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

func TestSubTest(t *testing.T) {
	t.Run("Muhammad", func(t *testing.T) {
		result := HelloWorld("Muhammad")
		require.Equal(t, "Hello Muhammad", result, "Result must be 'Hello Muhammad'")
	})
	t.Run("Arif", func(t *testing.T) {
		result := HelloWorld("Arif")
		require.Equal(t, "Hello Arif", result, "Result must be 'Hello Arif'")
	})
}

func TestMain(m *testing.M) {
	//before
	fmt.Println("BEFORE UNIT TEST")

	m.Run()

	//after
	fmt.Println("AFTER UNIT TEST")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Can not run on Macintosh")
	}

	result := HelloWorld("Arif")
	require.Equal(t, "Hello Arif", result, "Result must be 'Hello Arif'")

}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Arif")
	require.Equal(t, "Hello Arif", result, "Result must be 'Hello Arif'")
	fmt.Println("TestHelloWorld with Require Done")
}
func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Arif")
	assert.Equal(t, "Hello Arif", result, "Result must be 'Hello Arif'")
	fmt.Println("TestHelloWorld with Assert Done")
}

func TestHelloWorldArif(t *testing.T) {
	result := HelloWorld("Arif")

	if result != "Hello Arif" {
		//error
		// t.Fail()
		t.Error("Result must be Hello Arif")
	}

	fmt.Println("TestHelloWorldArif Done")

}

func TestHelloWorldJoko(t *testing.T) {
	result := HelloWorld("Joko")

	if result != "Hello Joko" {
		//error
		// t.FailNow()
		t.Fatal("Result must be Hello Joko")
	}

	fmt.Println("TestHelloWorldJoko Done")
}
