using System.Diagnostics;
using System.IO;
using System.Xml.Linq;
using static System.Net.WebRequestMethods;

namespace Glow
{
    class CommandExecutor
    {
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
            string cwd = Directory.GetCurrentDirectory();
            List<FileInfo> files = new List<FileInfo>();
            foreach (var item in Directory.GetFiles(cwd))
            {
                FileInfo file = new FileInfo(item);
                if (file.Name.StartsWith("0a--"))
                    continue;

                files.Add(file);
            }

            files.Sort((x, y) => x.Length.CompareTo(y.Length));
            files.Reverse();

            string compressedPath = Path.Combine(cwd, "compressed");

            if (!Directory.Exists(compressedPath))
                Directory.CreateDirectory(compressedPath);

            foreach (FileInfo file in files)
            {
                if (ExecuteCompressCommand(file.FullName, Path.Combine(compressedPath, file.Name)) == 0)
                    System.IO.File.Delete(file.FullName);
            }
        }
    }
}