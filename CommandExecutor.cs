using System.Diagnostics;
using System.IO;

namespace Glow {
    static class CommandExecutor
    {
        public static void Command(){



            Process p = new Process{
                StartInfo =
                 {
                     FileName = "git",                        
                     WorkingDirectory = @"C:\",
                     Arguments = "--version"
                 }
            };

            p.Start();
        }
    }
}