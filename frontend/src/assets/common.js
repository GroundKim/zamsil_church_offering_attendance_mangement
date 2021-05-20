export default function alertError (err) {
  let errStatusCode = err.response.status
  if (errStatusCode === 404) {
      alert(err.message + "데이터를 불러오는 도중 오류가 발생하였습니다 관리자에게 문의하십시오")
  }

  if (errStatusCode === 403) {
      alert("로그인을 해주세요")
      this.$router.push("/login")
  }
}