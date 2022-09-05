using Glow.Commands;
using System.Diagnostics;
using System.IO;

namespace Glow
{
    class VideoCompress
    {
        private VideoCompressOptions args;

        public VideoCompress(VideoCompressOptions args)
        {
            this.args = args;
        }

        private int ExecuteCompressCommand(string path, string destPath)
        {
            Process p = new Process
            {
                StartInfo =
                {
                    WorkingDirectory = Directory.GetCurrentDirectory(),
                    FileName = "ffmpeg",
                    Arguments = $"-i \"{path}\" -vcodec libx264 -crf 24 \"{destPath}\""
                }
            };

            p.Start();
            p.WaitForExit();
            return p.ExitCode;
        }

        public void CompressVideo()
        {
            string cwd; 
            if(args.Dir.Trim() == ".")
                cwd = Directory.GetCurrentDirectory();
            else 
                cwd = args.Dir;

            string[] whitelist = new string[] { ".mp4", ".mov", ".webm" };
            List<FileInfo> files = new List<FileInfo>();
            foreach (var item in Directory.GetFiles(cwd))
            {
                FileInfo file = new FileInfo(item);
                if (whitelist.Any(x => file.Name.EndsWith(x)))
                    files.Add(file);
            }

            if(files.Count == 0)
            {
                Console.WriteLine("No video found to compress in the given directory.");
                return;
            }

            files.Sort((x, y) => x.Length.CompareTo(y.Length));
            files.Reverse();

            string compressedPath = Path.Combine(cwd, "compressed");

            if (!Directory.Exists(compressedPath))
                Directory.CreateDirectory(compressedPath);

            foreach (FileInfo file in files)
            {
                if (ExecuteCompressCommand(file.FullName, Path.Combine(compressedPath, file.Name)) == 0)
                    File.Delete(file.FullName);
            }
        }
    }
}