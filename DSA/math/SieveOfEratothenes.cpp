#include <tr2/dynamic_bitset>
#include <bits/stdc++.h>
using namespace std;

// finds all prime numbers upto n
void SieveofEratosthenes1(int n)
{
    // consider index as a number
    vector<bool> prime(n + 1, true);
    prime[0] = prime[1] = false;
    for (int p = 2; p * p <= n; p++)
    {
        if (prime[p])
        {
            // multiple cutting
            for (int i = p * p; i <= n; i += p)
                prime[i] = false;
        }
    }
}
// finds all prime numbers upto n
void SieveofEratosthenes2(int n)
{

    tr2::dynamic_bitset<> prime(n + 1); // bitset is used to store boolean values in bits
    prime.set();                        // set all bits to 1
    prime[0] = prime[1] = false;

    // multiple of 2 are not prime
    for (int i = 4; i <= n; i += 2)
    {
        prime[i] = false;
    }

    // loop over odd number because all prime are odd
    for (int p = 3; p * p <= n; p += 2)
    {
        if (prime[p])
        {
            for (int i = p * p; i <= n; i += p)
            {
                prime[i] = false;
            }
        }
    }

    // print all primes
    for (int i = 0; i <= n; i++)
    {
        if (prime[i])
            cout << i << " ";
    }
}

int main()
{
    int n;
    cin >> n;
    SieveofEratosthenes2(n);
    return 0;
}