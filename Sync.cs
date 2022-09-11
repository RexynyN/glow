using Glow.Commands;
using System.Diagnostics;
using System.IO;




namespace Glow
{
    class CommandOutput{
        public int ExitCode { get; set; }
        public string Output { get; set; }
        public string Errors { get; set; }
    }

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
            foreach (string dir in GetLocalRepos())
            {
                Console.WriteLine($"Syncing down '{dir}'");
                SendCommand(dir, "git", "pull");
            }

            Console.WriteLine("All synced down.");
        }

        private void SyncUp()
        {
            foreach (string dir in GetLocalRepos())
            {
                Console.WriteLine($"Syncing up '{dir}'");
                SendCommand(dir, "git", "add .");
                CommandOutput response =  SendCommand(dir, "git", $"commit -m \"{CommitMessage()}\"");

                // This is the standard message when there`s no chances to commit
                // Change if there's any change to it.
                if (response.Output.Contains("nothing to commit, working tree clean")){
                    Console.WriteLine($"'{dir}' have nothing to commit, skipping...");
                    continue;
                }
                Console.WriteLine(response.Output);
                Console.WriteLine(response.Errors);



                response = SendCommand(dir, "git", "push");
                Console.WriteLine(response.Output);
                Console.WriteLine(response.Errors);

                if(response.ExitCode == 0)
                    Console.WriteLine($"'{dir}' have been synced up.");
                else
                    Console.WriteLine($"'{dir}' had an error in syncing up: " + response.Output);
            }

            Console.WriteLine("All synced up.");
        }

        private object CommitMessage()
        {
            return "Commit - " + DateTime.Now.ToString("dd/MM/yyyy HH:mm");
        }

        private IEnumerable<string> GetLocalRepos()
        {
            return File.ReadAllLines(Path.Combine(AppDomain.CurrentDomain.BaseDirectory, "sync.txt"));
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
                File.Create(filePath).Close();

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

        private CommandOutput SendCommand(string dir, string command, string arguments)
        {
            Process p = new Process
            {
                StartInfo =
                {
                    WorkingDirectory = args.Path,
                    FileName = command,
                    Arguments = arguments,
                    RedirectStandardOutput = true,
                    RedirectStandardError = true
                }
            };

            p.Start();
            p.WaitForExit();

            return new CommandOutput { ExitCode = p.ExitCode,  Output = p.StandardOutput.ReadToEnd(), Errors = p.StandardError.ReadToEnd()};
        }
    }
}