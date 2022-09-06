using Glow.Commands;
using System.Diagnostics;
using System.IO;


namespace Glow
{
    class Sync
    {
        private SyncOptions args;

        public Sync(SyncOptions args)
        {
            this.args = args;
        }

        public void RegisterRepo()
        {
            if(args.Path == ".")
                args.Path = Directory.GetCurrentDirectory();

            if(!Directory.Exists(args.Path) || !Directory.Exists(Path.Combine(args.Path, ".git")))
            {
                Console.WriteLine("The given directory doesn't exist or doesn't have a .git folder, aborting the registration.");
                return;
            }

            string list = File.ReadAllText(Path.Combine(AppDomain.CurrentDomain.BaseDirectory, "sync.txt"));
            list += args.Path;

            File.WriteAllText(Path.Combine(AppDomain.CurrentDomain.BaseDirectory, "sync.txt"), list);
            System.Console.WriteLine("Repository registered successfully");
        }

        public void GitAddCommand()
        {
            Process p = new Process
            {
                StartInfo =
                {
                    WorkingDirectory = @"C:\Users\017585631\Desktop\codes\Desafios",
                    FileName = "git",
                    Arguments = $"add .",
                    RedirectStandardOutput = true,
                    RedirectStandardError = true
                }
            };

            p.Start();
            p.WaitForExit();
        }

        public void GitCommitCommand()
        {
            Process p = new Process
            {
                StartInfo =
                {
                    WorkingDirectory = arg.Path,
                    FileName = "git",
                    Arguments = $"commit -m 'Clock'",
                    RedirectStandardOutput = true,
                    RedirectStandardError = true
                }
            };

            p.Start();

            string output = p.StandardOutput.ReadToEnd();
            if(output.Contains("nothing to commit, working tree clean"))
                System.Console.WriteLine("Seggs");

            p.WaitForExit();
        }

        

        public void OrchestrateCommand()
        {
            return;
        }
    }
}