using System;
using System.Collections.Generic;
using CommandLine;
using System.Linq;
using Glow.Commands;
using Glow;

// https://github.com/gsscoder/commandline/wiki/Latest-Version


public class Program
{
    public static void Main(string[] args)
    {
        Parser.Default.ParseArguments<PrimesOptions>(args)
          .WithParsed<PrimesOptions>(opts => new Primes(opts).PrimeFactory())
          .WithNotParsed(errs => HandleParseError(errs));

        ////var base64Img = new Base64Image 
        ////{
        ////    FileContents = File.ReadAllBytes(@"C:\Users\017585631\Desktop\codes\glow\ibmhangout.jpg"),
        ////    ContentType = "image/png"
        ////};

        ////string base64EncodedImg = base64Img.ToString();
        ////Console.WriteLine(base64EncodedImg);

        //Console.WriteLine("Hello, World!");

        //CommandExecutor.Command();

        //Console.WriteLine("End of the world");
        //Console.ReadKey();
    }

    //in case of errors or --help or --version
    static int HandleParseError(IEnumerable<Error> errs)
    {
        var result = -2;
        Console.WriteLine("errors {0}", errs.Count());
        if (errs.Any(x => x is HelpRequestedError || x is VersionRequestedError))
            result = -1;
        Console.WriteLine("Exit code {0}", result);
        return result;
    }
}
