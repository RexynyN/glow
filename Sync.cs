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

        private void CheckArgs()
        {
            if (args.Path == ".")
                args.Path = Directory.GetCurrentDirectory();
        }

        public void Start()
        {
            CheckArgs();

            switch (args.Action)
            {
                case "up":
                    SyncUp();
                    break;

                case "down":
                    SyncDown();
                    break;

                case "register":
                    RegisterRepo();
                    break;

                default:
                    break;
            }
        }

        public void SyncDown()
        {
            foreach (string dir in GetLocalDirs())
            {
                Console.WriteLine($"Syncing down '{dir}'");
                GitPullCommand(dir);
            }

            Console.WriteLine("All synced down.");
        }



        private IEnumerable<string> GetLocalDirs()
        {
            return File.ReadAllLines(Path.Combine(AppDomain.CurrentDomain.BaseDirectory, "sync.txt"));
        }

        private void SyncUp()
        {
            foreach (string dir in GetLocalDirs())
            {
                Console.WriteLine($"Syncing down '{dir}'");
                GitPullCommand(dir);
            }

            Console.WriteLine("All synced up.");
        }

        public void RegisterRepo()
        {
            if (!Directory.Exists(args.Path) || !Directory.Exists(Path.Combine(args.Path, ".git")))
            {
                Console.WriteLine("The given directory doesn't exist or doesn't have a .git folder, aborting the registration.");
                return;
            }

            string filePath = Path.Combine(AppDomain.CurrentDomain.BaseDirectory, "sync.txt");
            if (!File.Exists(filePath))
                File.Create(filePath);

            string list = "";
            try
            {
                list = File.ReadAllText(filePath);
            }
            catch (Exception)
            {
            }

            if(list.Contains(args.Path))
            {
                System.Console.WriteLine("Repository already registered, skipping...");
                return;
            }
            list += args.Path + "\n";

            File.WriteAllText(filePath, list);
            Console.WriteLine($"'{args.Path}' registered successfully");
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
                    WorkingDirectory = args.Path,
                    FileName = "git",
                    Arguments = $"commit -m 'Clock'",
                    RedirectStandardOutput = true,
                    RedirectStandardError = true
                }
            };

            p.Start();

            string output = p.StandardOutput.ReadToEnd();
            // This is the standard message when there`s no chances to commit
            // Change if there's any change to it.
            if (output.Contains("nothing to commit, working tree clean"))
                System.Console.WriteLine("Seggs");

            p.WaitForExit();
        }

        private void GitPullCommand(string path)
        {
            Process p = new Process
            {
                StartInfo =
                {
                    WorkingDirectory = path,
                    FileName = "git",
                    Arguments = $"pull",
                    RedirectStandardOutput = true,
                    RedirectStandardError = true
                }
            };

            p.Start();
            p.WaitForExit();
        }

        public void OrchestrateCommand()
        {
            return;
        }
    }
}