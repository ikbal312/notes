#include <bits/stdc++.h>
using namespace std;

int isPrime1(int n)
{
    // time:O(n)
    // space:O(1)
    // negative number is not prime
    if (n <= 1)
        return 0;

    for (int i = 2; i < n; i++)
    {
        if (n % i == 0)
            return 0;
    }
    return 1;
}
int isPrime2(int n)
{
    // time:O(sqrt(n))
    // space:O(1)

    if (n <= 1)
        return 0;
    // all divisors are less than sqrt(n), so we can check till sqrt(n+1) for precesion
    int limit = sqrt(n + 1);
    for (int i = 2; i < limit; i++) // i * i < n is same as i < sqrt(n)
    {
        if (n % i == 0)
            return 0;
    }
    return 1;
}
int isPrime3(int n)
{
    // time:O(sqrt(n))
    // space:O(1)
    // negative number is not prime
    if (n <= 1)
        return 0;

    // i * i < n is same as i < sqrt(n)
    for (int i = 2; i * i < n; i++)
    {
        // if n is divisible by i, then it is not prime
        if (n % i == 0)
            return 0;
    }
    return 1;
}

int main()
{
    int n;
    cin >> n;
    cout << isPrime1(n) << endl;
    cout << isPrime2(n) << endl;
    cout << isPrime3(n) << endl;
    return 0;
}