namespace cpp kuji
namespace java kuji
namespace php kuji

service KujiService {
    Response ThRegisterCandidatesWithKey(1: ReqCandidates req)
    Response ThPickOneByKey(1: ReqPickOneByKey req)
    Response ThPickOneByKeyAndIndex(1: ReqPickOneByKeyAndIndex req)
    Response ThPickAndDeleteOneByKey(1: ReqPickOneByKey req)
}

struct Response {
    1: i64 code,
    2: string message,
    3: Data data
}

struct Data {
    1: i64 id,
    2: string id_str,
}

struct ReqCandidates {
    1: string key,
    2: KujiCandidates candidates,
}

struct ReqPickOneByKey {
    1: string key,
}

struct ReqPickOneByKeyAndIndex {
    1: string key,
    2: i64 index,
}

struct KujiCandidate {
    1: i64 id,
    2: i64 weight,
}

struct KujiCandidates {
    1: list<KujiCandidate> candidates,
}