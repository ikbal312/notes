#include <bits/stdc++.h>
using namespace std;

// segmented sieve pseducode
// 1. find all prime number between 1 to sqrt(r)
// 2. find all prime number between l to r
// 3. print all prime number between l to r
// time complexity: O(nlog(log(n)))
// space complexity: O(sqrt(n))

void segmentedSieve1(int l, int r)
{

    // prime between 1 to sqrt(r)
    int limit = sqrt(r) + 1;
    vector<bool> basePrime(limit, true);
    basePrime[0] = basePrime[1] = false;
    for (int p = 2; p * p <= limit; p++)
    {
        if (basePrime[p])
        {
            for (int i = p * p; i <= limit; i += p)
                basePrime[i] = false;
        }
    }

    // prime between l to r
    int range = r - l + 1;
    vector<bool> prime(range, true);
    // even is not prime
    for (int p = l; p <= r; p++)
    {
        if (p % 2 == 0)
        {
            prime[p - l] = false;
        }
    }

    // all odd number between l to r beacuse all prime are odd
    for (int i = 3; i <= limit; i += 2)
    {
        if (basePrime[i]) // if prime
        {

            for (int j = i * i; j <= r; j += i)
            {
                prime[j - l] = false;
            }
        }
    }

    for (int i = 0; i < range; i++)
    {
        if (prime[i])
        {
            cout << i + l << " ";
        }
    }
}

//

int main()
{

    int l, r;
    cin >> l >> r;
    segmentedSieve1(l, r);
    return 0;
}