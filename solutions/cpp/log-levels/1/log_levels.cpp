#include <string>

namespace log_line {
std::string message(std::string line) {
    return line.substr(line.find("]:") + 3);
}

std::string log_level(std::string line) {
    const int start = line.find("[");
    const int end = line.find("]:");

    return line.substr(start + 1, end - start - 1);
}

std::string reformat(std::string line) {
    return log_line::message(line) + " (" + log_line::log_level(line) + ")";
}
}  // namespace log_line
