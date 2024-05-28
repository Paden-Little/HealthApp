import axios from "axios";
const bcrypt = require('bcryptjs');
const fixedSalt = "$2a$10$1234567890123456789012";

const hashPassword = async (password: string) => {
  return await bcrypt.hash(password, fixedSalt);
}


export function useAuth() {
  const registerPatient = async (patient: any) => {
    patient.password = await hashPassword(patient.password);
    return await axios.post("127.0.0.1/patient", patient);
  }

  const registerProvider = async (provider: any) => {
    provider.password = await hashPassword(provider.password);
    return await axios.post("127.0.0.1/provider", provider);
  }

  const loginPatient = async (patient: any) => {
    patient.password = await hashPassword(patient.password);
    return await axios.post("127.0.0.1/patient/login", patient);
  }

  const loginProvider = async (provider: any) => {
    provider.password = await hashPassword(provider.password);
    return await axios.post("127.0.0.1/provider/login", provider);
  }

  return { registerPatient, registerProvider, loginPatient, loginProvider };
}