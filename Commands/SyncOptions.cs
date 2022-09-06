using CommandLine.Text;
using CommandLine;

namespace Glow.Commands
{
    [Verb("sync", HelpText = "Sync eveything you need, from github repos to... well, github repos.")]
    class SyncOptions
    {
        [Value(0, Required = true, HelpText = "Sync up, sync down, register.")]
        public string Action { get; set; }

        [Option('p', "path", Default = ".", HelpText = "Set the output to a .txt file.", SetName = "register")]
        public string Path { get; set; }
    }
}
