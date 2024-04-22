import { defineStore } from 'pinia'

import { postUsersLogin } from '@/helpers/api/api'
import { setToken } from '@/helpers/funcs/auth-utils'

export const useUserStore = defineStore('user', () => {
  async function login(email: string, password: string) {
    try {
      const response = await postUsersLogin({ email, password })
      if (response.status === "OK") {
        setToken(response.token);
      }
    } catch (error) {
      console.error(error)
    }
  }

  function base64ToArrayBuffer(base64: string) {
    const binaryString = atob(base64);
    const bytes = new Uint8Array(binaryString.length);
    for (let i = 0; i < binaryString.length; i++) {
        bytes[i] = binaryString.charCodeAt(i);
    }
    return bytes.buffer;
}

  function bufferToBase64URL(buffer: ArrayBuffer) {
    const bytes = new Uint8Array(buffer);
    let string = '';
    bytes.forEach(b => string += String.fromCharCode(b));

    const base64 = btoa(string);
    return base64.replace(/\+/g, '-').replace(/\//g, '_').replace(/=/g, '');
  }

  async function webAuthnSignUp(email: string) {
    try {
      const response = await fetch('/api/v1/users/authn/signup/begin', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({ email })
      
      });
      const data = await response.json();
      console.log(data)
      const pubKey = data.options.publicKey      ;
      console.log(pubKey);
      const publicKey = {
        challenge: base64ToArrayBuffer(pubKey.challenge),
        rp: {
          id: pubKey.rp.id,
          name: pubKey.rp.name
        },
        user: {
          id: pubKey.user.id,
          name: email,
          displayName: email
        },
        pubKeyCredParams: pubKey.pubKeyCredParams,
        timeout: pubKey.timeout,
        attestation: "none" as AttestationConveyancePreference,
      };
      const cred = await navigator.credentials.create({ publicKey });
      if (cred) {
        const credential = {
          id: cred.id,
          type: cred.type,
          response: {
            attestationObject: bufferToBase64URL(cred.response.attestationObject),
            clientDataJSON: bufferToBase64URL(cred.response.clientDataJSON)
          }
        
        };
        localStorage.setItem('credential', JSON.stringify(credential));
        const response = await fetch('/api/v1/users/authn/signup/finish', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json'
          },
          body: JSON.stringify({ email })
        
        });
        console.log(response);
      }
    } catch (error) {
      console.error(error)
    }
  }

  return {
    login,
    webAuthnSignUp
  }
})
