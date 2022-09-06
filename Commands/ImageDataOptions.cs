using CommandLine.Text;
using CommandLine;

namespace Glow.Commands
{
    [Verb("imagedata", HelpText = "Transform images into base 64 data uris and vice-versa.")]
    class ImageDataOptions
    {
        [Value(0, Required = true, Default=".", HelpText = "Location of the file/directory.")]
        public string Location { get; set; }

        [Option('d', "dir", Default = false, HelpText = "Signifies that the location is a directory of images.")]
        public bool Dir { get; set; }

        [Option('m', "max", Default = 100, HelpText = "Max number of images to fit in a single file.")]
        public int Max { get; set; }

        [Option('c', "convert", Default = false, HelpText = "Convert images to base 64 datauri.")]
        public bool Convert { get; set; }

        [Option('r', "revert", Default = false, HelpText = "Revert base 64 datauri to image files.")]
        public bool Revert { get; set; }
    }
}
