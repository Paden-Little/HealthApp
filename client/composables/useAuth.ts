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

  async function loginPatient(patient: Login): Promise<boolean | undefined> {
    return $fetch("/api/patient/login", {
      method: "POST",
      body: JSON.stringify(patient),
    }).then((res) => {
      const loginResponse: LoginResponse = res as LoginResponse;
      pid.value = loginResponse.id;
      token.value = loginResponse.token;
      return true; 
    }).catch((err) => {
      console.log(err);
      return false;
    });
  }

  async function loginProvider(provider: Login): Promise<boolean | undefined> {
    return $fetch("/api/provider/login", {
      method: "POST",
      body: JSON.stringify(provider),
    }).then((res) => {
      const loginResponse: LoginResponse = res as LoginResponse;
      pid.value = loginResponse.id;
      token.value = loginResponse.token;
      return true; 
    }).catch((err) => {
      console.log(err);
      return false; 
    });
  }

  async function getPatientData(): Promise<Patient | null> {
    return $fetch(`/api/patient/${pid.value}`, {
      headers: {
        Authorization: `Bearer ${token.value}`,
      },
    }).then((res) => {
      return res as Patient;
    }).catch((err) => {
      console.log(err);
      return null;
    })
  }

  async function getProviderData(): Promise<Provider | null> {
    return $fetch(`/api/provider/${pid.value}`, {
      headers: {
        Authorization: `Bearer ${token.value}`,
      },
    }).then((res) => {
      return res as Provider;
    }).catch((err) => {
      console.log(err);
      return null;
    })
  }

  async function logoutUser() {
    pid.value = "";
    token.value = "";
  }
  
  return { registerPatient, registerProvider, loginPatient, loginProvider, getPatientData, getProviderData, logoutUser };
}