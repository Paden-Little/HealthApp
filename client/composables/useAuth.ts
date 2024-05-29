import bcrypt from 'bcryptjs';
const fixedSalt = '$2a$10$1234567890123456789012';
interface LoginResponse {
  id: string;
  token: string;
}

// cookies
const pid = useCookie("pid");
const token = useCookie("token");
const type = useCookie("type");

// state
const user = useState<Patient | Provider | undefined>();

const hashPassword = async (password: string) => {
  return await bcrypt.hash(password, fixedSalt);
};

export function useAuth() {
  async function registerPatient(patient: Patient) {
    patient.password = await hashPassword(patient.password || '');
    await $fetch('/api/patient', {
      method: 'POST',
      body: JSON.stringify(patient),
    })
      .catch(err => {
        console.log(err);
        return false;
      });
    await loginPatient({ email: patient.email, password: patient.password }).then((res) => {
      if (res) {
        return true;
      }
    }).catch((err) => {
      console.log(err);
      return false;
    });
  }

  async function registerProvider(provider: Provider) {
    provider.password = await hashPassword(provider.password || '');
    await $fetch('/api/provider', {
      method: 'POST',
      body: JSON.stringify(provider),
    })
      .catch(err => {
        console.log(err);
        return false;
      });
    await loginPatient({ email: provider.email, password: provider.password }).then((res) => {
      if (res) {
        return true;
      }
    }).catch((err) => {
      console.log(err);
      return false;
    });
  }

  async function loginPatient(patient: Login): Promise<boolean | undefined> {
    return $fetch('/api/patient/login', {
      method: 'POST',
      body: JSON.stringify(patient),
    }).then((res) => {
      const loginResponse: LoginResponse = res as LoginResponse;
      pid.value = loginResponse.id;
      token.value = loginResponse.token;
      type.value = "patient";
      getPatientData().then((data) => {
        user.value = data as Patient;
      })
      return true;
    }).catch((err) => {
      console.log(err);
      return false;
    });
  }

  async function loginProvider(provider: Login): Promise<boolean | undefined> {
    return $fetch('/api/provider/login', {
      method: 'POST',
      body: JSON.stringify(provider),
    }).then((res) => {
      const loginResponse: LoginResponse = res as LoginResponse;
      pid.value = loginResponse.id;
      token.value = loginResponse.token;
      type.value = "provider";
      getProviderData().then((data) => {
        user.value = data as Provider;
      })
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
    })
      .then(res => {
        return res as Patient;
      })
      .catch(err => {
        console.log(err);
        return null;
      });
  }

  async function getProviderData(): Promise<Provider | null> {
    return $fetch(`/api/provider/${pid.value}`, {
      headers: {
        Authorization: `Bearer ${token.value}`,
      },
    })
      .then(res => {
        return res as Provider;
      })
      .catch(err => {
        console.log(err);
        return null;
      });
  }

  async function logoutUser() {
    type.value = undefined;
    pid.value = undefined;
    token.value = undefined;
    user.value = undefined;
  }

  async function getUserAppointment() {
    return $fetch(`/api/appointment/${pid.value}`, {
      headers: {
        Authorization: `Bearer ${token.value}`,
      },
    }).then((res) => {
      return res as Appointment[];
    }).catch((err) => {
      console.log(err);
      return [];
    })
  }

  return { registerPatient, registerProvider, loginPatient, loginProvider, getPatientData, getProviderData, logoutUser, getUserAppointment, user };
}
