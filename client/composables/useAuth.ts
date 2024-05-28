const bcrypt = require('bcryptjs');
const fixedSalt = "$2a$10$1234567890123456789012";

const hashPassword = async (password: string) => {
  return await bcrypt.hash(password, fixedSalt);
}


export function useAuth() {
  const registerPatient = async (patient: any) => {
    patient.password = await hashPassword(patient.password);
    let { data } = useFetch("/api/patient", {
      method: "POST",
      body: JSON.stringify(patient),
    })
    return data;
  }

  const registerProvider = async (provider: any) => {
    provider.password = await hashPassword(provider.password);
    let { data } = useFetch("/api/provider", {
      method: "POST",
      body: JSON.stringify(provider),
    })
    return data
  }

  const loginPatient = async (patient: any) => {
    patient.password = await hashPassword(patient.password);
    let { data } = useFetch("/api/provider", {
      method: "POST",
      body: JSON.stringify(patient),
    })
    return data
  }

  const loginProvider = async (provider: any) => {
    provider.password = await hashPassword(provider.password);
    let { data } = useFetch("/api/provider", {
      method: "POST",
      body: JSON.stringify(provider),
    })
    return data
  }

  const getPatientData = async (token: string, pid: string) => {
    let { data } = useFetch(`/api/patient/${pid}`, {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    })
    return data;
  }

  const getProviderData = async (token: string, pid: string) => {
    let { data } = useFetch(`/api/provider/${pid}`, {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    })
    return data;
  }

  return { registerPatient, registerProvider, loginPatient, loginProvider, getPatientData, getProviderData};
}