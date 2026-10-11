export function maskEmail(email: string) {
  return email.replace(/(.).*(@.*)/, "$1***$2");
}
