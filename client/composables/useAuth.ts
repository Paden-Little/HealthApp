import bcrypt from "bcryptjs";
const fixedSalt = "$2a$10$1234567890123456789012";

const hashPassword = async (password: string) => {
  return await bcrypt.hash(password, fixedSalt);
}


export function useAuth() {
  const registerPatient = async (patient: Patient) => {
    patient.password = await hashPassword(patient.password || '');
    let { data } = useFetch("/api/patient", {
      method: "POST",
      body: JSON.stringify(patient),
    })
  }

  const registerProvider = async (provider: Provider) => {
    provider.password = await hashPassword(provider.password || '');
    let { data } = useFetch("/api/provider", {
      method: "POST",
      body: JSON.stringify(provider),
    })
  }

  const loginPatient = async (patient: Login) => {
    patient.password = await hashPassword(patient.password);
    let { data } = useFetch("/api/patient", {
      method: "POST",
      body: JSON.stringify(patient),
    })
    // Set cookies
    console.log(data)
  }

  const loginProvider = async (provider: Login) => {
    let pid = useCookie("pid");
    let token = useCookie("token");
    provider.password = await hashPassword(provider.password);
    let { data } = useFetch("/api/provider", {
      method: "POST",
      body: JSON.stringify(provider),
    })
    console.log(data)
  }

  const getPatientData = async () => {
    let pid = useCookie("pid");
    let token = useCookie("token");
    let { data } = useFetch(`/api/patient/${pid.value}`, {
      headers: {
        Authorization: `Bearer ${token.value}`,
      },
    })
    return data;
  }

  const getProviderData = async () => {
    let pid = useCookie("pid");
    let token = useCookie("token");
    let { data } = useFetch(`/api/provider/${pid.value}`, {
      headers: {
        Authorization: `Bearer ${token.value}`,
      },
    })
    return data;
  }

  return { registerPatient, registerProvider, loginPatient, loginProvider, getPatientData, getProviderData };
}