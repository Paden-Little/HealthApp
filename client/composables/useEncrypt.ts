export function useEncrypt() {
  const bcrypt = require('bcryptjs');

  const hashPassword = async (password: string) => {
    const salt = await bcrypt.genSalt(10);
    return await bcrypt.hash(password, salt);
  }

  const comparePasswords = async (password: string, hashedPassword: string) => {
    return await bcrypt.compare(password, hashedPassword);
  }

  return {hashPassword, comparePasswords};
}