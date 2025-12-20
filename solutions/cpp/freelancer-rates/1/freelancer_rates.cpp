// INFO: Headers from the standard library should be inserted at the top via
#include <math.h>

// daily_rate calculates the daily rate given an hourly rate
double daily_rate(double hourly_rate) {
    const int DAILY_TIME = 8;

    return DAILY_TIME * hourly_rate;
}

// apply_discount calculates the price after a discount
double apply_discount(double before_discount, double discount) {
    return before_discount - (before_discount * discount / 100);
}

// monthly_rate calculates the monthly rate, given an hourly rate and a discount
// The returned monthly rate is rounded up to the nearest integer.
int monthly_rate(double hourly_rate, double discount) {
    const int MONTH_DAYS = 22;
    const double rounded_hourly_rate = (hourly_rate - (int)hourly_rate);
    const double rounded_value = (rounded_hourly_rate < 0.5 && rounded_hourly_rate != 0.0) ? 0.5 : 0.0;

    return round(MONTH_DAYS * apply_discount(daily_rate(hourly_rate), discount) + rounded_value);
}

// days_in_budget calculates the number of workdays given a budget, hourly rate,
// and discount The returned number of days is rounded down (take the floor) to
// the next integer.
int days_in_budget(int budget, double hourly_rate, double discount) {
    return budget / apply_discount(daily_rate(hourly_rate), discount);
}
