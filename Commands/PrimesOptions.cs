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
        [Value(0, Required=true, Default=100, HelpText="The number of primes to return.")]
        public int Threshold { get; set; }
    }
}
