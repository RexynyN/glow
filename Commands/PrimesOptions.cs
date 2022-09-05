using CommandLine.Text;
using CommandLine;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Glow.Commands
{
    [Verb("primes", HelpText = "Get the first n-th prime numbers, for whatever reason")]
    class PrimesOptions
    {
        [Value(0, Required = true, Default = 100, HelpText = "The number of primes to return.")]
        public int Threshold { get; set; }

        [Option('t', "txt", Default = false, HelpText = "Set the output to a .txt file.")]
        public bool Txt { get; set; }

        [Option('j', "json", Default = false, HelpText = "Set the output to a .json file.")]
        public bool Json { get; set; }
    }
}
