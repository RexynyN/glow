using Glow.Commands;
using System.Text.Json;
using System.Xml.Linq;

namespace Glow
{
    class Primes
    {
        private PrimesOptions args;

        public Primes(PrimesOptions args)
        {
            this.args = args;
        }

        public void PrimeFactory()
        {
            int number = 3;
            List<int> primes = new List<int>(args.Threshold) { 2 };
            int i = 1;
            bool guard;
            while (i < args.Threshold)
            {
                guard = true;
                double root = Math.Sqrt(number);
                foreach (var prime in primes)
                {
                    if (prime > root)
                        break;

                    if (number % prime == 0)
                    {
                        guard = false;
                        break;
                    }
                }

                if (guard)
                {
                    primes.Add(number);
                    i += 1;
                }
                number++;
            }

            if (args.Json)
            {
                var json = JsonSerializer.Serialize(primes);
                File.WriteAllText(Path.Combine(Directory.GetCurrentDirectory(), "primes.json"), json);
            }
            else if (args.Txt)
            {
                string output = string.Join(" ", primes.ToArray());
                File.WriteAllText(Path.Combine(Directory.GetCurrentDirectory(), "primes.txt"), output);
            }
            else
            {
                foreach (var prime in primes)
                    Console.Write(prime + " ");
            }
        }
    }
}
