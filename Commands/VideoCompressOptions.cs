using CommandLine.Text;
using CommandLine;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Glow.Commands
{
    [Verb("videocompress", HelpText = "Compress videos using ffmpeg, but in a nice one-line command.")]
    class VideoCompressOptions
    {
        [Value(0, Required = false, Default=".", HelpText = "The number of primes to return.")]
        public string Dir { get; set; }
    }
}
