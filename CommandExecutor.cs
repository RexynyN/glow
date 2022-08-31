using System.Diagnostics;

namespace Glow {
    static class CommandExecutor
    {
        public static void Command(){
             Process p = new Process{
                StartInfo =
                 {
                     FileName = "node",                        
                     WorkingDirectory = @"C:\",
                     Arguments = "--version"
                 }
            };

            p.Start();
        }
    }
}