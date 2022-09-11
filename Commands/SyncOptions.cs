using CommandLine.Text;
using CommandLine;

namespace Glow.Commands
{
    [Verb("sync", HelpText = "Sync eveything you need, from github repos to... well, github repos.")]
    class SyncOptions
    {
        [Value(0, Required = true, HelpText = "Sync up, sync down, register.")]
        public string ?Action { get; set; }

        [Option('p', "path", Default = ".", HelpText = "The path to the repository.", SetName = "register")]
        public string ?Path { get; set; }

        [Option('v', "verbose", Default = false, HelpText = "Set the command to verbose mode.", SetName = "register")]
        public bool Verbose{ get; set; }
        
        [Option('l', "list", Default = false, HelpText = "Set the command to verbose mode.", SetName = "register")]
        public bool List{ get; set; }
    }
}
