import axios from "axios";

export function useAuth() {
  const registerPatient = async (patient: any) => {
    return await axios.post("127.0.0.1/patient", patient);
  }

  const registerProvider = async (provider: any) => {
    return await axios.post("127.0.0.1/provider", provider);
  }

  const loginPatient = async (patient: any) => {
    return await axios.post("127.0.0.1/patient/login", patient);
  }

  const loginProvider = async (provider: any) => {
    return await axios.post("127.0.0.1/provider/login", provider);
  }

  return {registerPatient, registerProvider, loginPatient, loginProvider};
}