function b64enc(buf) {
    return base64js.fromByteArray(buf)
                   .replace(/\+/g, "-")
                   .replace(/\//g, "_")
                   .replace(/=/g, "");
}

function b64RawEnc(buf) {
    return base64js.fromByteArray(buf)
    .replace(/\+/g, "-")
    .replace(/\//g, "_");
}

function hexEncode(buf) {
    return Array.from(buf)
                .map(function(x) {
                    return ("0" + x.toString(16)).substr(-2);
				})
                .join("");
}

// https://stackoverflow.com/questions/10730362/get-cookie-by-name
function getCookie(name) {
    var nameEQ = name + "=";
    var ca = document.cookie.split(';');
    for(var i = 0; i < ca.length; i++) {
        var c = ca[i];
        while (c.charAt(0)==' ') {
            c = c.substring(1,c.length);
        }
        if (c.indexOf(nameEQ) == 0) {
            return c.substring(nameEQ.length,c.length);
        }
    }
    return null;
}

async function fetch_json(url, options) {
    const response = await fetch(url, options);    
    const body = await response.json();
    if (body.fail)
        throw body.fail;
    return body;
}

/**
 * REGISTRATION FUNCTIONS
 */

/**
 * Callback after the registration form is submitted.
 * @param {Event} e 
 */
const registrationListenerHelper = async (form_id, begin_url, finish_url, e) => {
    e.preventDefault();

    // Gather the data in the form
    const form = document.querySelector(form_id);
    const formData = new FormData(form);

    // POST the data to the server to generate the PublicKeyCredentialCreateOptions
    let credentialCreateOptionsFromServer;
    try {
        credentialCreateOptionsFromServer = await getCredentialCreateOptionsFromServer(formData, begin_url);
    } catch (err) {
        alert("Failed to generate credential request options: " + err);
        window.location.reload(false);
        return;
    }

    // Convert certain members of the PublicKeyCredentialCreateOptions into
    // byte arrays as expected by the spec.
    const publicKeyCredentialCreateOptions = transformCredentialCreateOptions(credentialCreateOptionsFromServer);
    
    // Request the authenticator(s) to create a new credential keypair.
    let credential;
    try {
        credential = await navigator.credentials.create({
            publicKey: publicKeyCredentialCreateOptions
        });
    } catch (err) {
        alert("Error creating credential: " + err);
        window.location.reload(false);
        return;
    }

    // We now have a new credential! We now need to encode the byte arrays
    // in the credential into strings, for posting to our server.
    const newAssertionForServer = transformNewAssertionForServer(credential);

    // POST the transformed credential data to the server for validation
    // and storing the public key
    let assertionValidationResponse;
    try {
        assertionValidationResponse = await postNewAssertionToServer(formData, newAssertionForServer, finish_url);
    } catch (err) {
        alert("Server validation of credential failed: " + err);
        window.location.reload(false);
        return;
    }
    
    console.warn("Redirecting to: " + assertionValidationResponse.nexturl);
    window.location.assign(assertionValidationResponse.nexturl);
}

const createRegistrationListener = (form_id, begin_url, finish_url) => {
    async function listener_fn(e) {
        return registrationListenerHelper(form_id, begin_url, finish_url, e);
    }

    return listener_fn;
}

/**
 * Get PublicKeyCredentialRequestOptions for this user from the server
 * formData of the registration form
 * @param {FormData} formData 
 */
const getCredentialRequestOptionsFromServer = async (formData, begin_url) => {
    return await fetch_json(
        begin_url,
        {
            method: "POST",
            body: formData
        }
    );
}

const transformCredentialRequestOptions = (credentialRequestOptionsFromServer) => {
    let {challenge, allowCredentials} = credentialRequestOptionsFromServer;

    challenge = Uint8Array.from(
        atob(challenge.replace(/\_/g, "/").replace(/\-/g, "+")), c => c.charCodeAt(0));

    allowCredentials = allowCredentials.map(credentialDescriptor => {
        let {id} = credentialDescriptor;
        id = id.replace(/\_/g, "/").replace(/\-/g, "+");
        id = Uint8Array.from(atob(id), c => c.charCodeAt(0));
        return Object.assign({}, credentialDescriptor, {id});
    });

    const transformedCredentialRequestOptions = Object.assign(
        {},
        credentialRequestOptionsFromServer,
        {challenge, allowCredentials});

    return transformedCredentialRequestOptions;
};


/**
 * Get PublicKeyCredentialRequestOptions for this user from the server
 * formData of the registration form
 * @param {FormData} formData 
 */
const getCredentialCreateOptionsFromServer = async (formData, begin_url) => {
    return await fetch_json(
        begin_url,
        {
            method: "POST",
            body: formData
        }
    );
}

/**
 * Transforms items in the credentialCreateOptions generated on the server
 * into byte arrays expected by the navigator.credentials.create() call
 * @param {Object} credentialCreateOptionsFromServer 
 */
const transformCredentialCreateOptions = (credentialCreateOptionsFromServer) => {
    let {challenge, user} = credentialCreateOptionsFromServer;
    user.id = Uint8Array.from(
        atob(credentialCreateOptionsFromServer.user.id
            .replace(/\_/g, "/")
            .replace(/\-/g, "+")
            ), 
        c => c.charCodeAt(0));

    challenge = Uint8Array.from(
        atob(credentialCreateOptionsFromServer.challenge
            .replace(/\_/g, "/")
            .replace(/\-/g, "+")
            ),
        c => c.charCodeAt(0));
    
    const transformedCredentialCreateOptions = Object.assign(
            {}, credentialCreateOptionsFromServer,
            {challenge, user});

    return transformedCredentialCreateOptions;
}

/**
 * AUTHENTICATION FUNCTIONS
 */


/**
 * Callback build blocks to be executed after submitting login form
 * @param {Event} e 
 */
const attestationBegin_URL = async (form_id, begin_url) => {
    // Gather the data in the form
    const form = document.querySelector(form_id);
    const formData = new FormData(form);

    // POST the login data to the server to retrieve the `PublicKeyCredentialRequestOptions`
    const credentialRequestOptionsFromServer = await getCredentialRequestOptionsFromServer(formData, begin_url);
    return credentialRequestOptionsFromServer;
}

const attestationBegin_Cookie = async (begin_src) => {
    // TODO: Return `null` if cookie is not found
    const credentialRequestOptionsFromServer = JSON.parse(decodeURIComponent(getCookie(begin_src)));
    return credentialRequestOptionsFromServer
}

const attestationBegin_FormField = async (form_id, field_name) => {
    // Gather the data in the form
    const form = document.querySelector(form_id);
    const formData = new FormData(form);

    if (formData.get(field_name) === "") {
        return null;
    }

    const credentialRequestOptionsFromServer = JSON.parse(formData.get(field_name));
    return credentialRequestOptionsFromServer
}

const attestationFinish_URL = async (credentialRequestOptionsFromServer, finish_url, form_id) => {
    let formData;
    if (form_id !== null) {
        // Gather the data in the form
        const form = document.querySelector(form_id);
        formData = new FormData(form);
    } else {
        formData = new FormData();
    }

    let response;

    // Webauthn is enabled
    if (credentialRequestOptionsFromServer) {
        // Convert certain members of the PublicKeyCredentialRequestOptions into
        // byte arrays as expected by the spec.    
        const transformedCredentialRequestOptions = transformCredentialRequestOptions(
            credentialRequestOptionsFromServer);

        // Request the authenticator to create an assertion signature using the
        // credential private key
        const assertion = await navigator.credentials.get({
            publicKey: transformedCredentialRequestOptions,
        });

        // We now have an authentication assertion! encode the byte arrays contained
        // in the assertion data as strings for posting to the server
        const transformedAssertionForServer = transformAssertionForServer(assertion);

        // POST the assertion to the server for verification.
        response = await postAssertionToServer(transformedAssertionForServer, finish_url, formData);
    } else {
        // Perform a non-Webauthn POST
        response = await postAssertionToServer("", finish_url, formData);
    }

    // Go to the url in the `response`
    window.location.assign(response.url);
}

/**
 * Transforms the binary data in the credential into base64 strings
 * for posting to the server.
 * @param {PublicKeyCredential} newAssertion 
 */
const transformNewAssertionForServer = (newAssertion) => {
    const attObj = new Uint8Array(
        newAssertion.response.attestationObject);
    const clientDataJSON = new Uint8Array(
        newAssertion.response.clientDataJSON);
    const rawId = new Uint8Array(
        newAssertion.rawId);

    return {
        id: newAssertion.id,
        rawId: b64enc(rawId),
        type: newAssertion.type,
        response: {
            attestationObject: b64enc(attObj),
            clientDataJSON: b64enc(clientDataJSON),
        },
    };
}

/**
 * Posts the new credential data to the server for validation and storage.
 * @param {Object} credentialDataForServer 
 */
const postNewAssertionToServer = async (formData, credentialDataForServer, finish_url) => {
    return await fetch_json(
        finish_url, 
        {
            method: "POST",
            headers: 
            {
                'X-CSRF-TOKEN': formData.get('_csrf')
            },
            body: JSON.stringify(credentialDataForServer)
        });
}

/**
 * Encodes the binary data in the assertion into strings for posting to the server.
 * @param {PublicKeyCredential} newAssertion 
 */
const transformAssertionForServer = (newAssertion) => {
    const authData = new Uint8Array(newAssertion.response.authenticatorData);
    const clientDataJSON = new Uint8Array(newAssertion.response.clientDataJSON);
    const rawId = new Uint8Array(newAssertion.rawId);
    const sig = new Uint8Array(newAssertion.response.signature);

    return {
        id: newAssertion.id,
        rawId: b64enc(rawId),
        type: newAssertion.type,
        response: {
            authenticatorData: b64enc(authData),
            clientDataJSON: b64enc(clientDataJSON),
            signature: b64enc(sig),
        }
    };
};

/**
 * Post the assertion to the server for validation and logging the user in. 
 * @param {Object} assertionDataForServer 
 */
const postAssertionToServer = async (assertionDataForServer, finish_url, formData) => {
    // Pass over the webauthn assertion in JSON format
    formData.set('webauthn_data', JSON.stringify(assertionDataForServer));

    return await fetch(
        finish_url,
        {
            method: "POST",
            body: formData
        }
    );
}
