using Glow.Commands;
using SixLabors.ImageSharp;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Glow
{
    class Primes
    {
        private PrimesOptions Args;

        public Primes(PrimesOptions args)
        {
            Args = args;
        }


        public int  PrimeFactory()
        {
            int number = 3;
            List<int> primes = new List<int>(Args.Threshold);
            int i = 1;
            bool guard = false;
            while (i < Args.Threshold)
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

            foreach (var prime in primes)
                Console.Write(prime + " ");

            return 0;
        }
    }
}
