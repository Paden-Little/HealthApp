import bcrypt from "bcryptjs";
const fixedSalt = "$2a$10$1234567890123456789012";

interface LoginResponse {
  id: string;
  token: string;
}

const pid = useCookie("pid");
const token = useCookie("token");

const hashPassword = async (password: string) => {
  return await bcrypt.hash(password, fixedSalt);
}

export function useAuth() {
  const registerPatient = async (patient: Patient) => {
    // patient.password = await hashPassword(patient.password || '');
    // let { data } = useFetch("/api/patient", {
    //   method: "POST",
    //   body: JSON.stringify(patient),
    // })
  }

  const registerProvider = async (provider: Provider) => {
    // provider.password = await hashPassword(provider.password || '');
    // let { data } = useFetch("/api/provider", {
    //   method: "POST",
    //   body: JSON.stringify(provider),
    // })
  }

  const loginPatient = async (patient: Login) => {
    $fetch("/api/patient/login", {
      method: "POST",
      body: JSON.stringify(patient),
    }).then((res) => {
      const loginResponse: LoginResponse = res as LoginResponse;
      pid.value = loginResponse.id;
      token.value = loginResponse.token;
      return true;
    })
    return false;
  };

  const loginProvider = async (provider: Login) => {
    $fetch("/api/provider/login", {
      method: "POST",
      body: JSON.stringify(provider),
    }).then((res) => {
      const loginResponse: LoginResponse = res as LoginResponse;
      pid.value = loginResponse.id;
      token.value = loginResponse.token;
      return true;
    })
    return false;
  }

  const getPatientData = async () => {
    $fetch(`/api/patient/${pid.value}`, {
      headers: {
        Authorization: `Bearer ${token.value}`,
      },
    }).then((res) => {
      return res as Patient;
    })
    return null;
  }

  const getProviderData = async () => {
    $fetch(`/api/patient/${pid.value}`, {
      headers: {
        Authorization: `Bearer ${token.value}`,
      },
    }).then((res) => {
      return res as Patient;
    })
    return null;
  }

  return { registerPatient, registerProvider, loginPatient, loginProvider, getPatientData, getProviderData };
}